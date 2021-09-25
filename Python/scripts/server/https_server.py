"""
@version: 3.9
@title: https 服务端程序

在http.server的基础上 使用SSLContext进行包装

建议使用 SSLContext 实例的 SSLContext.wrap_socket() 来将套接字包装为 SSLSocket 对象。 
辅助函数 create_default_context() 会返回一个新的带有安全默认设置的上下文。

包装一个现有的 Python 套接字 sock 并返回一个 SSLContext.sslsocket_class 的实例 
(默认为 SSLSocket)。 
返回的 SSL 套接字会绑定上下文、设置以及证书。 
sock 必须是一个 SOCK_STREAM 套接字；其他套接字类型不被支持

单向认证流程： 1.客户端say hello服务端 2.服务端将证书、公钥等发给客户端 3.客户端CA验证证书，成功继续、不成功弹出选择页面 4.客户端告知服务端所支持的加密算法 5.服务端选择最高级别加密算法明文通知客户端 6.客户端生成随机对称密匙key，使用服务端公钥加密发送给服务端 7.服务端使用私钥解密，获取对称密匙key 8.后续客户端与服务端使用该密匙key进行加密通信

双向认证流程(后续通讯所用的对称key由两边一起生成)： 1.客户端say hello服务端 2.服务端将证书、公钥等发给客户端 3.客户端CA验证证书，成功继续、不成功弹出选择页面 4.客户端将自己的证书和公钥发送给服务端 5.服务端验证客户端证书，如不通过直接断开连接 6.客户端告知服务端所支持的加密算法 7.服务端选择最高级别加密算法使用客户端公钥加密后发送给客户端 8.客户端收到后使用私钥解密并生成随机对称密匙key，使用服务端公钥加密发送给服务端 9.服务端使用私钥解密，获取对称密匙key 10.后续客户端与服务端使用该密匙key进行加密通信

- load_cert_chain: 证书以及私钥 context.load_cert_chain("server-cert.pem","server-key.pem")


创建证书的步骤
ca-key.pem ca-req.csr ca-cert.pem server-key.pem server-req.csr server-cert.pem

# ca 证书颁发者
- 创建私钥：openssl genrsa -out ca-key.pem 1024
- 创建csr证书请求: openssl req -new -key ca-key.pem -out ca-req.csr -subj "/C=CN/ST=BJ/L=BJ/O=BJ/OU=BJ/CN=BJ"
- 生成crt证书: openssl x509 -req -in ca-req.csr -out ca-cert.pem -signkey ca-key.pem -days 3650

# 服务端私钥以及证书
- 创建服务端私钥: openssl genrsa -out server-key.pem 1024
- 创建csr证书: openssl req -new -out server-req.csr -key server-key.pem -subj "/C=CN/ST=BJ/L=BJ/O=BJ/OU=BJ/CN=BJ"
- 生成crt证书: openssl x509 -req -in server-req.csr -out server-cert.pem -signkey server-key.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -days 3650

# 客户端私钥以及证书(用于双向验证)
- 确认证书: openssl verify -CAfile ca-cert.pem  server-cert.pem

需要注意的是如果使用浏览器访问必须勾选信任该证书才能够继续使用
"""
from typing import Type, Tuple
import json
import ssl
from http.server import BaseHTTPRequestHandler, HTTPServer

from scripts.server.cert import path


class _HTTPHandler(BaseHTTPRequestHandler):
    TODO = ["You get the response!!"]

    def do_GET(self):
        if self.path != "/":
            self.send_error(404, "Page not Found!")
            return

        resp = json.dumps(self.TODO)
        self.send_response(200, message="OK")
        self.send_header("Content-Type", "application/json")
        self.end_headers()
        self.wfile.write(resp.encode())


class HTTPSServer(HTTPServer):
    def __init__(
        self,
        address: Tuple[str, int],
        handler: Type[BaseHTTPRequestHandler] = _HTTPHandler,
    ):
        super().__init__(address, handler)

        self.ssl_context = ssl.SSLContext(ssl.PROTOCOL_TLS_SERVER)
        self.ssl_context.load_cert_chain(
            path("server-cert.pem"), path("server-key.pem")
        )
        self.socket = self.ssl_context.wrap_socket(self.socket, server_side=True)


if __name__ == "__main__":
    if len(argv := __import__("sys").argv) > 1 and argv[1] == "test":
        import doctest

        doctest.testmod()

    if len(argv) > 1 and argv[1] == "start":
        server = HTTPSServer(("localhost", 8080))
        try:
            server.serve_forever()
        finally:
            server.server_close()