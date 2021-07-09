from typing import Any, Optional
from pydantic import BaseModel


class Message(BaseModel):
    code: int
    message: str
    data: Any