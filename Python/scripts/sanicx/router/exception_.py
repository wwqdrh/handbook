import json
from typing import Any, Sequence

from pydantic import ValidationError, create_model
from pydantic.error_wrappers import ErrorList

from sanicx.router.encoders import jsonable_encoder

RequestErrorModel = create_model("Request")
WebSocketErrorModel = create_model("WebSocket")


class RequestValidationError(ValidationError):
    def __init__(self, errors: Sequence[ErrorList], *, body: Any = None) -> None:
        self.body = body
        super().__init__(errors, RequestErrorModel)

    @property
    def pretty_errors(self) -> str:
        return json.dumps(jsonable_encoder(self.errors()))
