"""
一些配置类的schema
"""

from typing import Optional, TypeVar, Generic, Union, Any, Literal, List

from pydantic import BaseModel, Field, validator, ValidationError
from starlette.datastructures import Secret


class DatabaseURL(BaseModel):
    """
    Database 配置类
    """

    class Config:
        arbitrary_types_allowed = True  # 允许任意类型, 这里是为了适配Secret属性
        # allow_population_by_alias = True

    drivername: Literal["mysql", "sqlite", "postgresql"] = Field(
        ..., description="The database driver., mysql、sqlite、postgresql"
    )
    host: str = Field("localhost", description="Server host.")
    port: Optional[Union[str, int]] = Field(None, description="Server access port.")
    username: Optional[str] = Field(None, description="Username")
    password: Optional[Union[str, Secret]] = Field(None, description="Password")
    database: str = Field(
        ...,
        description="Database name.",
    )

    @validator("drivername", always=True)
    def validate_drivername(cls, v: str) -> str:
        if v not in ("mysql", "sqlite", "postgresql"):
            raise ValueError("只能选择 mysql sqlite postgresql三种")
        return v

    @property
    def url(self) -> str:
        if self.drivername == "mysql":
            return "mysql://{username}:{password}@{host}:{port}/{database}".format(
                username=self.username,
                password=self.password,
                host=self.host,
                port=self.port,
                database=self.database,
            )

        return ""
