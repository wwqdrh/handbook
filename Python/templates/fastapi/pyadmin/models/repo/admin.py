from typing import Union
from tortoise.fields.data import (
    CharField,
    IntEnumField,
    DatetimeField,
    IntField,
    SmallIntField,
)

from pyadmin.utils.models import Model
from ._spec import UsedStatus, DeleteStatus
from pyadmin.config import SALT
from pyadmin.utils.hash import get_hash_password


class Admin(Model):
    class Meta:
        table = "admin"

    username = CharField(max_length=255, unique=True)
    password = CharField(max_length=255)
    nickname = CharField(max_length=255)
    mobile = CharField(max_length=20)
    is_used = IntEnumField(UsedStatus)
    is_deleted = IntEnumField(DeleteStatus)
    created_at = DatetimeField(auto_now_add=True)
    created_user = CharField(max_length=255)
    updated_at = DatetimeField(auto_now=True)
    updated_user = CharField(max_length=255)

    @classmethod
    async def check_login(
        cls, username: str, password: str
    ) -> Union[Exception, "Admin"]:
        if (row := await cls.filter(username=username).first()) is None:
            raise Exception("用户不存在")

        if get_hash_password(password, row.salt) != row.password:
            raise Exception("密码错误")

        return row


class AdminMenu(Model):
    class Meta:
        table = "admin_menu"

    admin_id = IntField(max_length=11, index=True)
    menu_id = IntField(max_length=11)
    created_at = DatetimeField(auto_now_add=True)
    created_user = CharField(max_length=255)


class Authorized(Model):
    class Meta:
        table = "authorized"

    business_key = CharField(max_length=32)
    business_secret = CharField(max_length=60)
    business_developer = CharField(max_length=60)
    remark = CharField(max_length=255)
    is_used = IntEnumField(UsedStatus)
    is_deleted = IntEnumField(DeleteStatus)
    created_at = DatetimeField(auto_now_add=True)
    created_user = CharField(max_length=255)
    updated_at = DatetimeField(auto_now=True)
    updated_user = CharField(max_length=255)


class AuthorizedApi(Model):
    class Meta:
        table = "authorized_api"

    business_key = CharField(max_length=32)
    method = CharField(max_length=30)
    api = CharField(max_length=100)
    is_deleted = IntEnumField(DeleteStatus)
    created_at = DatetimeField(auto_now_add=True)
    created_user = CharField(max_length=255)
    updated_at = DatetimeField(auto_now=True)
    updated_user = CharField(max_length=255)


class Menu(Model):
    class Meta:
        table = "menu"

    pid = IntField(max_length=11)
    name = CharField(max_length=32)
    link = CharField(max_length=100)
    icon = CharField(max_length=60)
    level = SmallIntField()
    sort = SmallIntField()
    is_used = IntEnumField(UsedStatus)
    is_deleted = IntEnumField(DeleteStatus)
    created_at = DatetimeField(auto_now_add=True)
    created_user = CharField(max_length=255)
    updated_at = DatetimeField(auto_now=True)
    updated_user = CharField(max_length=255)


class MenuAction(Model):
    class Meta:
        table = "menu_action"

    menu_id = IntField()
    method = CharField(max_length=30)
    api = CharField(max_length=100)
    is_deleted = IntEnumField(DeleteStatus)
    created_at = DatetimeField(auto_now_add=True)
    created_user = CharField(max_length=255)
    updated_at = DatetimeField(auto_now=True)
    updated_user = CharField(max_length=255)
