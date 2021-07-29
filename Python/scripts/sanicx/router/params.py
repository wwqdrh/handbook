from enum import Enum
from typing import Any, Optional, Callable

from pydantic.fields import FieldInfo


class ParamTypes(Enum):
    query = "query"
    header = "header"
    path = "path"
    cookie = "cookie"


class Param(FieldInfo):
    in_: ParamTypes

    def __init__(
        self,
        default: Any,
        *,
        alias: Optional[str] = None,
        title: Optional[str] = None,
        description: Optional[str] = None,
        gt: Optional[float] = None,
        ge: Optional[float] = None,
        lt: Optional[float] = None,
        le: Optional[float] = None,
        min_length: Optional[int] = None,
        max_length: Optional[int] = None,
        regex: Optional[str] = None,
        deprecated: Optional[bool] = None,
        **extra: Any,
    ):
        self.deprecated = deprecated
        super().__init__(
            default,
            alias=alias,
            title=title,
            description=description,
            gt=gt,
            ge=ge,
            lt=lt,
            le=le,
            min_length=min_length,
            max_length=max_length,
            regex=regex,
            **extra,
        )

    def __repr__(self) -> str:
        return f"{self.__class__.__name__}({self.default})"


class FlaskParam(Param):
    """
    view_func中默认的参数，由flask传递参数进来
    """

    in_ = ParamTypes.path

    def __init__(
        self,
        default: Any,
        *,
        alias: Optional[str] = None,
        title: Optional[str] = None,
        description: Optional[str] = None,
        gt: Optional[float] = None,
        ge: Optional[float] = None,
        lt: Optional[float] = None,
        le: Optional[float] = None,
        min_length: Optional[int] = None,
        max_length: Optional[int] = None,
        regex: Optional[str] = None,
        deprecated: Optional[bool] = None,
        **extra: Any,
    ):
        super().__init__(
            default,
            alias=alias,
            title=title,
            description=description,
            gt=gt,
            ge=ge,
            lt=lt,
            le=le,
            min_length=min_length,
            max_length=max_length,
            regex=regex,
            deprecated=deprecated,
            **extra,
        )


class Query(Param):
    in_ = ParamTypes.query

    def __init__(
        self,
        default: Any,
        *,
        alias: Optional[str] = None,
        title: Optional[str] = None,
        description: Optional[str] = None,
        gt: Optional[float] = None,
        ge: Optional[float] = None,
        lt: Optional[float] = None,
        le: Optional[float] = None,
        min_length: Optional[int] = None,
        max_length: Optional[int] = None,
        regex: Optional[str] = None,
        deprecated: Optional[bool] = None,
        **extra: Any,
    ):
        super().__init__(
            default,
            alias=alias,
            title=title,
            description=description,
            gt=gt,
            ge=ge,
            lt=lt,
            le=le,
            min_length=min_length,
            max_length=max_length,
            regex=regex,
            deprecated=deprecated,
            **extra,
        )


class Header(Param):
    in_ = ParamTypes.header

    def __init__(
        self,
        default: Any,
        *,
        alias: Optional[str] = None,
        convert_underscores: bool = True,
        title: Optional[str] = None,
        description: Optional[str] = None,
        gt: Optional[float] = None,
        ge: Optional[float] = None,
        lt: Optional[float] = None,
        le: Optional[float] = None,
        min_length: Optional[int] = None,
        max_length: Optional[int] = None,
        regex: Optional[str] = None,
        deprecated: Optional[bool] = None,
        **extra: Any,
    ):
        self.convert_underscores = convert_underscores
        super().__init__(
            default,
            alias=alias,
            title=title,
            description=description,
            gt=gt,
            ge=ge,
            lt=lt,
            le=le,
            min_length=min_length,
            max_length=max_length,
            regex=regex,
            deprecated=deprecated,
            **extra,
        )


class Cookie(Param):
    in_ = ParamTypes.cookie

    def __init__(
        self,
        default: Any,
        *,
        alias: Optional[str] = None,
        title: Optional[str] = None,
        description: Optional[str] = None,
        gt: Optional[float] = None,
        ge: Optional[float] = None,
        lt: Optional[float] = None,
        le: Optional[float] = None,
        min_length: Optional[int] = None,
        max_length: Optional[int] = None,
        regex: Optional[str] = None,
        deprecated: Optional[bool] = None,
        **extra: Any,
    ):
        super().__init__(
            default,
            alias=alias,
            title=title,
            description=description,
            gt=gt,
            ge=ge,
            lt=lt,
            le=le,
            min_length=min_length,
            max_length=max_length,
            regex=regex,
            deprecated=deprecated,
            **extra,
        )


class Body(FieldInfo):
    def __init__(
        self,
        default: Any,
        *,
        embed: bool = False,
        validator: Optional[Callable] = None,
        error_message: Optional[str] = None,
        media_type: str = "application/json",
        alias: Optional[str] = None,
        title: Optional[str] = None,
        description: Optional[str] = None,
        gt: Optional[float] = None,
        ge: Optional[float] = None,
        lt: Optional[float] = None,
        le: Optional[float] = None,
        min_length: Optional[int] = None,
        max_length: Optional[int] = None,
        regex: Optional[str] = None,
        **extra: Any,
    ):
        self.embed = embed
        self.media_type = media_type
        self.validator = validator
        self.error_message = error_message
        super().__init__(
            default,
            alias=alias,
            title=title,
            description=description,
            gt=gt,
            ge=ge,
            lt=lt,
            le=le,
            min_length=min_length,
            max_length=max_length,
            regex=regex,
            **extra,
        )

    def __repr__(self) -> str:
        return f"{self.__class__.__name__}({self.default})"


class Form(Body):
    def __init__(
        self,
        default: Any,
        *,
        validator: Optional[Callable] = None,
        error_message: Optional[str] = None,
        media_type: str = "application/x-www-form-urlencoded",
        alias: Optional[str] = None,
        title: Optional[str] = None,
        description: Optional[str] = None,
        gt: Optional[float] = None,
        ge: Optional[float] = None,
        lt: Optional[float] = None,
        le: Optional[float] = None,
        min_length: Optional[int] = None,
        max_length: Optional[int] = None,
        regex: Optional[str] = None,
        **extra: Any,
    ):
        super().__init__(
            default,
            embed=True,
            validator=validator,
            error_message=error_message,
            media_type=media_type,
            alias=alias,
            title=title,
            description=description,
            gt=gt,
            ge=ge,
            lt=lt,
            le=le,
            min_length=min_length,
            max_length=max_length,
            regex=regex,
            **extra,
        )


class File(Form):
    def __init__(
        self,
        default: Any,
        *,
        validator: Optional[Callable] = None,
        error_message: Optional[str] = None,
        media_type: str = "multipart/form-data",
        alias: Optional[str] = None,
        title: Optional[str] = None,
        description: Optional[str] = None,
        gt: Optional[float] = None,
        ge: Optional[float] = None,
        lt: Optional[float] = None,
        le: Optional[float] = None,
        min_length: Optional[int] = None,
        max_length: Optional[int] = None,
        regex: Optional[str] = None,
        **extra: Any,
    ):
        super().__init__(
            default,
            validator=validator,
            error_message=error_message,
            media_type=media_type,
            alias=alias,
            title=title,
            description=description,
            gt=gt,
            ge=ge,
            lt=lt,
            le=le,
            min_length=min_length,
            max_length=max_length,
            regex=regex,
            **extra,
        )
