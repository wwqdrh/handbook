from typing import Optional
from datetime import datetime, timedelta

from fastapi import Depends, Request, HTTPException, status
from fastapi.security import (
    OAuth2PasswordBearer,
    SecurityScopes,
)
from pydantic import BaseModel, ValidationError
from jose import JWTError, jwt

from ._base import credentials_exception, no_permission_exception
from pyadmin.utils.hash import get_hash_password
from pyadmin.config import SECRET_KEY, ALGORITHM, ACCESS_TOKEN_EXPIRE_MINUTES
from pyadmin.models.repo.admin import Admin


__all__ = ("is_login", "TokenData", "encode_token", "decode_token")


oauth2_scheme = OAuth2PasswordBearer(tokenUrl="/admin/login")
# scopes={"user": "a normal user", "staff": "a staff", "admin": "a admin"},


class TokenData(BaseModel):
    username: str
    nickname: str
    mobile: str
    is_used: int
    is_deleted: int


def encode_token(admin: Admin) -> str:
    """传入admin用户模型生成对应token"""
    return jwt.encode(
        {
            "username": admin.username,
            "nickname": admin.nickname,
            "mobile": admin.mobile,
            "is_used": admin.is_used.value,
            "is_deleted": admin.is_deleted.value,
            "exp": datetime.utcnow() + timedelta(minutes=ACCESS_TOKEN_EXPIRE_MINUTES),
        },
        SECRET_KEY,
        algorithm=ALGORITHM,
    )


def decode_token(security_scopes: SecurityScopes, token: str) -> TokenData:
    """解码token"""
    authenticate_value = (
        f'Bearer scope="{security_scopes.scope_str}"'
        if security_scopes.scopes
        else "Bearer"
    )
    try:
        payload = jwt.decode(token, SECRET_KEY, algorithms=[ALGORITHM])
        token_data = TokenData(
            username=payload.get("username"),
            nickname=payload.get("nickname"),
            mobile=payload.get("mobile"),
            is_used=payload.get("is_used"),
            is_deleted=payload.get("is_deleted"),
        )
    except (JWTError, ValidationError):
        raise credentials_exception(authenticate_value)

    return token_data


async def is_login(
    request: Request,
    security_scopes: SecurityScopes,
    token: str = Depends(oauth2_scheme),
) -> TokenData:
    """
    判断是否登录
    """
    token_data = decode_token(security_scopes, token)

    # TODO: 查询数据库判断用户是否存在
    return token_data
