from typing import Union

from fastapi import Depends, Security
from fastapi.security import OAuth2PasswordRequestForm

from pyadmin.apps.api.v1._base import api_v1
from pyadmin.apps.api.v1._schema import AuthToken
from pyadmin.context import Application
from pyadmin.apps.schema import Message
from pyadmin.apps.depend import app
from pyadmin.apps.auth.admin_auth import is_login, TokenData, encode_token
from pyadmin.models.repo.admin import Admin


@api_v1.post(
    "/admin/login",
    tags=["api"],
    description="aouth2验证机制，获取token",
    response_model=Union[Message, AuthToken],
)
async def admin_login(
    form_data: OAuth2PasswordRequestForm = Depends(),
):
    """
    form_data: username、password
    """
    admin = await Admin.check_login(form_data.username, form_data.password)

    return {
        "access_token": encode_token(admin),
        "token_type": "bearer",
    }


@api_v1.get(
    "/hello",
    tags=["api"],
    response_model=Message,
)
async def admin_hello(
    app: Application = Depends(app),
    admin: TokenData = Security(is_login),
):
    return {"code": 2000, "message": "world"}
