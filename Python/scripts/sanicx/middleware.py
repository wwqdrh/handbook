import typing as T

from sanic.request import Request
from sanic.response import BaseHTTPResponse

if T.TYPE_CHECKING:
    from sanicx.app import SanicX


@T.runtime_checkable
class IMiddleware(T.Protocol):
    def register(self, app: "SanicX", handle_name: T.Tuple[str], attach_to: str):
        if attach_to == "request":
            app.register_named_middleware(
                self.request_middleware, handle_name, attach_to=attach_to
            )
        elif attach_to == "response":
            app.register_named_middleware(
                self.response_middleware, handle_name, attach_to=attach_to
            )

    def request_middleware(self, request: Request):
        raise NotImplementedError

    def response_middleware(self, request: Request, response: BaseHTTPResponse):
        raise NotImplementedError
