import typing as T
from hashlib import md5
from random import Random

from sanic import response
from sanic.request import Request
from sanic.response import BaseHTTPResponse

from sanicx.middleware import IMiddleware
from sanicx.utils.auth.httpauth_compat import parse_authorization_header
from sanicx.utils.auth.httpauth_compat import safe_str_cmp, Authorization

__all__ = ("HTTPBasicAuth", "HTTPTokenAuth", "HTTPDigestAuth")


class IHTTPAuth(IMiddleware, T.Protocol):
    scheme: str
    realm: str
    get_password_cb: T.Callable = None
    auth_error_cb: T.Callable = None

    get_password = lambda self, fn: setattr(self, "get_password_cb", fn)
    error_handler = lambda self, fn: setattr(self, "auth_error_cb", fn)

    @staticmethod
    def get_username(request: Request):
        if not request.ctx.authorization:
            return ""
        return request.ctx.authorization.username

    def get_auth(self, request) -> T.Optional[Authorization]:
        """
        获取用户信息
        """
        auth = parse_authorization_header(request.headers.get("Authorization"))
        try:
            if auth is None and "Authorization" in request.headers:
                auth_headers = request.headers["Authorization"]
                auth_type, value = auth_headers.split(None, 1)
                auth = Authorization(auth_type, {"token": value})
        except ValueError:
            # The Authorization header is either empty or has no token
            pass

        # if the auth type does not match, we act as if there is no auth
        # this is better than failing directly, as it allows the callback
        # to handle special cases, like supporting multiple auth types
        if auth is not None and auth.type.lower() != self.scheme.lower():
            auth = None

        return auth

    def get_auth_password(self, auth: Authorization) -> T.Any:
        """
        获取用户密码
        """
        if self.get_password_cb is None:
            return None

        password = None
        if auth and auth.username:
            password = self.get_password_cb(auth.username)

        return password

    def auth_failed_callback(self, request: Request):
        """
        验证失败的回调
        """
        if self.auth_error_cb is None:
            res = response.text("Unauthorized Access", 401)
        else:
            res = self.auth_error_cb(request)
        if res.status == 200:
            res.status = 401
        if "WWW-Authenticate" not in res.headers.keys():
            res.headers["WWW-Authenticate"] = self.authenticate_header(request)
        return res

    def request_middleware(self, request: Request):
        """
        auth验证中间件
        """
        auth = self.get_auth(request)
        request.ctx.authorization = auth
        if request.method != "OPTIONS":  # pragma: no cover
            password = self.get_auth_password(auth)
            if not self.authenticate(request, auth, password):
                return self.auth_failed_callback(request)

    def response_middleware(self, request: Request, response: BaseHTTPResponse):
        pass

    def authenticate_header(self, request: Request):
        raise NotImplementedError

    def authenticate(self, request: Request, auth: Authorization, stored_password: str):
        raise NotImplementedError


class HTTPBasicAuth(IHTTPAuth):
    hash_password = lambda self, fn: setattr(self, "hash_password_callback", fn)
    verify_password = lambda self, fn: setattr(self, "verify_password_callback", fn)

    def __init__(self, scheme: str = "Basic", realm: str = None):
        self.scheme = scheme
        self.realm = realm or "Authentication Required"
        self.hash_password_callback = None
        self.verify_password_callback = None

    def authenticate_header(self, request: Request):
        return '{0} realm="{1}"'.format(self.scheme, self.realm)

    def authenticate(self, request: Request, auth: Authorization, stored_password: str):
        if auth:
            username = auth.username
            client_password = auth.password
        else:
            username = ""
            client_password = ""
        if self.verify_password_callback:
            return self.verify_password_callback(username, client_password)
        if not auth:
            return False
        if self.hash_password_callback:
            try:
                client_password = self.hash_password_callback(client_password)
            except TypeError:
                client_password = self.hash_password_callback(username, client_password)
        return (
            client_password is not None
            and stored_password is not None
            and safe_str_cmp(client_password, stored_password)
        )


class HTTPDigestAuth(IHTTPAuth):
    def __init__(
        self, scheme: str = "Digest", realm: str = None, use_ha1_pw: bool = False
    ):
        self.scheme = scheme
        self.realm = realm or "Authentication Required"
        self.use_ha1_pw = use_ha1_pw
        self._random = lambda: md5(str(Random().random()).encode("utf-8")).hexdigest()

        def default_generate_nonce(request):
            request.ctx.session["auth_nonce"] = self._random()
            return request.ctx.session["auth_nonce"]

        def default_verify_nonce(request, nonce):
            session_nonce = request.ctx.session.get("auth_nonce")
            if nonce is None or session_nonce is None:
                return False
            return safe_str_cmp(nonce, session_nonce)

        def default_generate_opaque(request):
            request.ctx.session["auth_opaque"] = self._random()
            return request.ctx.session["auth_opaque"]

        def default_verify_opaque(request, opaque):
            session_opaque = request.ctx.session.get("auth_opaque")
            if opaque is None or session_opaque is None:
                return False
            return safe_str_cmp(opaque, session_opaque)

        self.generate_nonce_callback = default_generate_nonce
        self.verify_nonce_callback = default_verify_nonce
        self.generate_opaque_callback = default_generate_opaque
        self.verify_opaque_callback = default_verify_opaque

    def set_generate_nonce(self, f):
        self.generate_nonce_callback = f
        return f

    def set_generate_opaque(self, f):
        self.generate_opaque_callback = f
        return f

    def set_verify_nonce(self, f):
        self.verify_nonce_callback = f
        return f

    def set_verify_opaque(self, f):
        self.verify_opaque_callback = f
        return f

    def generate_ha1(self, username, password):
        a1 = username + ":" + self.realm + ":" + password
        a1 = a1.encode("utf-8")
        return md5(a1).hexdigest()

    def authenticate_header(self, request: Request):
        nonce = self.generate_nonce_callback(request)
        opaque = self.generate_opaque_callback(request)
        return '{0} realm="{1}",nonce="{2}",opaque="{3}"'.format(
            self.scheme, self.realm, nonce, opaque
        )

    def authenticate(self, request: Request, auth: Authorization, stored_password: str):
        if not (
            auth
            and auth.username
            and auth.realm
            and auth.uri
            and auth.nonce
            and auth.response
            and stored_password
        ):
            return False
        if not (
            self.verify_nonce_callback(request, auth.nonce)
            and self.verify_opaque_callback(request, auth.opaque)
        ):
            return False
        if self.use_ha1_pw:
            ha1 = stored_password
        else:
            a1 = ":".join([auth.username, auth.realm, stored_password])
            ha1 = md5(a1.encode("utf-8")).hexdigest()
        a2 = ":".join([request.method, auth.uri])
        ha2 = md5(a2.encode("utf-8")).hexdigest()
        a3 = ":".join([ha1, auth.nonce, ha2])
        return safe_str_cmp(md5(a3.encode("utf-8")).hexdigest(), auth.response)


class HTTPTokenAuth(IHTTPAuth):
    def __init__(self, scheme="Bearer", realm=None):
        self.scheme = scheme
        self.realm = realm or "Authentication Required"
        self.verify_token_callback = None

    def set_verify_token(self, f):
        self.verify_token_callback = f
        return f

    def token(self, request):
        if not request.ctx.authorization:
            return ""
        return request.ctx.authorization.get("token")

    def authenticate_header(self, request: Request):
        pass

    def authenticate(self, request, auth, stored_password):
        if auth:
            token = auth["token"]
        else:
            token = ""
        if self.verify_token_callback:
            return self.verify_token_callback(token)
        return False
