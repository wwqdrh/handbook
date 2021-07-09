from typing import Any, Optional, Type

from pydantic import BaseModel


class AuthToken(BaseModel):
    access_token: str
    token_type: str
    expires_in: Optional[int]
    refresh_token: Optional[str]
    example_parameter: Optional[str]