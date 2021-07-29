"""
email插件
"""
from flask import current_app
import blinker
import smtplib
import time
from email.utils import formatdate, formataddr, make_msgid, parseaddr
from email.header import Header
from email import policy


__all__ = ("Email", "EmailContext", "Message", "Connection")

BadHeaderError = type("BadHeaderError", (object,), {})
string_types = (str,)
text_type = str
message_policy = policy.SMTP

signals = blinker.Namespace()

email_dispatched = signals.signal(
    "email-dispatched",
    doc=(
        "Signal sent when an email is dispatched. This signal will also be sent "
        "in testing mode, even though the email will not actually be sent."
    ),
)


class Message:
    """connection 传递消息时的中间对象

    :attr sender: 发送人
    :attr recipients: 收件人
    :attr bcc
    :attr cc
    """

    def __init__(
        self,
        subject="",
        recipients=None,
        body=None,
        html=None,
        sender=None,
        cc=None,
        bcc=None,
        attachments=None,
        reply_to=None,
        date=None,
        charset=None,
        extra_headers=None,
        mail_options=None,
        rcpt_options=None,
    ):

        sender = sender or current_app.extensions["mail"].default_sender

        if isinstance(sender, tuple):
            sender = "%s <%s>" % sender

        self.recipients = recipients or []
        self.subject = subject
        self.sender = sender
        self.reply_to = reply_to
        self.cc = cc or []
        self.bcc = bcc or []
        self.body = body
        self.html = html
        self.date = date
        self.msgId = make_msgid()
        self.charset = charset
        self.extra_headers = extra_headers
        self.mail_options = mail_options or []
        self.rcpt_options = rcpt_options or []
        self.attachments = attachments or []

    @property
    def send_to(self):
        return set(self.recipients) | set(self.bcc or ()) | set(self.cc or ())

    @classmethod
    def sanitize_address(self, addr: str, encoding="utf-8"):
        nm, addr = parseaddr(addr)

        try:
            nm = Header(nm, encoding).encode()
        except UnicodeEncodeError:
            nm = Header(nm, "utf-8").encode()
        try:
            addr.encode("ascii")
        except UnicodeEncodeError:  # IDN
            if "@" in addr:
                localpart, domain = addr.split("@", 1)
                localpart = str(Header(localpart, encoding))
                domain = domain.encode("idna").decode("ascii")
                addr = "@".join([localpart, domain])
            else:
                addr = Header(addr, encoding).encode()
        return formataddr((nm, addr))

    def has_bad_headers(self):
        pass

    def as_bytes(self):
        pass


class Connection:
    """管理邮件服务相关的连接

    :attr mail: EmailContext实例，包含需要发送的邮箱的相关数据
    :attr host: SMTP_SSL 或者 SMTP服务
    :attr num_emails: 邮箱数量相关
    """

    def __init__(self, mail: "EmailContext"):
        self.mail = mail
        self.host = None
        self.num_emails = 0

    def __enter__(self):
        if not self.mail.suppress:
            self.host = self.configure_host
        self.num_emails = 0
        return self

    def __exit__(self, exc_type, exc_val, exc_trace):
        if self.host is not None:
            self.host.quit()

    @property
    def configure_host(self):
        """相关smtp_ssl smtp连接"""
        if self.mail.use_ssl:
            host = smtplib.SMTP_SSL(self.mail.server, self.mail.port)
        else:
            host = smtplib.SMTP(self.mail.server, self.mail.port)

        host.set_debuglevel(int(self.mail.debug))

        if self.mail.use_tls:
            host.starttls()
        if self.mail.username and self.mail.password:
            host.login(self.mail.username, self.mail.password)

        return host

    def send(self, message: "Message", envelope_from=None):
        """验证以及发送消息

        :param message: Message的实例
        :param envelope_from: Email address to be used in MAIL FROM command.
        """
        assert message.send_to, "No recipients have been added"

        assert message.sender, (
            "The message does not specify a sender and a default sender "
            "has not been configured"
        )

        if message.has_bad_headers():
            raise BadHeaderError

        if message.date is None:
            message.date = time.time()

        if self.host:
            self.host.sendmail(
                message.sanitize_address(envelope_from or message.sender),
                list(message.sanitize_addresses(message.send_to)),
                message.as_bytes(),
                message.mail_options,
                message.rcpt_options,
            )

        email_dispatched.send(message, app=current_app._get_current_object())

        self.num_emails += 1

        if self.num_emails == self.mail.max_emails:
            self.num_emails = 0
            if self.host:
                self.host.quit()
                self.host = self.configure_host()


class EmailContext:
    """实例会作为flask的插件，也就是说flask使用email服务的时候是对这个进行处理"""

    def __init__(
        self,
        server,
        username,
        password,
        port,
        use_tls,
        use_ssl,
        default_sender,
        debug,
        max_emails,
        suppress,
        ascii_attachments=False,
    ):
        self.server = server
        self.username = username
        self.password = password
        self.port = port
        self.use_tls = use_tls
        self.use_ssl = use_ssl
        self.default_sender = default_sender
        self.debug = debug
        self.max_emails = max_emails
        self.suppress = suppress
        self.ascii_attachments = ascii_attachments

    @property
    def connection(self):
        # 不直接使用self构建connection，因为必须保证emaiil注册到了flask应用上
        try:
            email_context = current_app.extensions["mail"]
        except (AttributeError, KeyError):
            raise RuntimeError(
                "The curent application was not configured with Mail.init_app(app)"
            )
        else:
            return Connection(email_context)

    def send(self, message: "Message"):
        """
        :param message: a Message instance.
        """
        with self.connection as connection:
            connection.send(message)


class Email:
    """EmailContext的代理类，用于帮助注册到flask的extension上"""

    def __init__(self, app=None):
        """
        :param app: Flask instance
        """
        self.app = app
        if app is not None:
            self.email_context = self.init_app(app)
        else:
            self.email_context = None

    def __getattr__(self, name):
        return getattr(self.email_context, name, None)

    def init_app(self, app):
        """
        flask 插件的通用格式，init_app
        :param app: Flask application instance
        """
        if getattr(self, "email_context", None) is not None:  # 已经被实例化过了
            return
        self.email_context = EmailContext(
            app.config.get("MAIL_SERVER", "127.0.0.1"),
            app.config.get("MAIL_USERNAME"),
            app.config.get("MAIL_PASSWORD"),
            app.config.get("MAIL_PORT", 25),
            app.config.get("MAIL_USE_TLS", False),
            app.config.get("MAIL_USE_SSL", False),
            app.config.get("MAIL_DEFAULT_SENDER"),
            int(app.config.get("MAIL_DEBUG", app.debug)),
            app.config.get("MAIL_MAX_EMAILS"),
            app.config.get("MAIL_SUPPRESS_SEND", app.testing),
            app.config.get("MAIL_ASCII_ATTACHMENTS", False),
        )

        # register extension with app
        app.extensions = getattr(app, "extensions", {})
        app.extensions["mail"] = self.email_context
