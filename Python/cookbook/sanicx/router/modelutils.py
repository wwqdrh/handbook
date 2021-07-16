import functools
import inspect
from copy import deepcopy
from typing import (
    Any,
    Callable,
    cast,
    Dict,
    List,
    Mapping,
    Optional,
    Sequence,
    Tuple,
    Type,
    Union,
)

from dependency_injector.wiring import Provide, Provider
from pydantic import BaseConfig
from pydantic import BaseModel, validator
from pydantic.class_validators import Validator
from pydantic.error_wrappers import ErrorWrapper
from pydantic.errors import MissingError
from pydantic.fields import (
    FieldInfo,
    ModelField,
    Required,
    SHAPE_LIST,
    SHAPE_SEQUENCE,
    SHAPE_SET,
    SHAPE_SINGLETON,
    SHAPE_TUPLE,
    SHAPE_TUPLE_ELLIPSIS,
    UndefinedType,
)
from pydantic.schema import get_annotation_from_field_info
from pydantic.typing import evaluate_forwardref, ForwardRef
from pydantic.utils import lenient_issubclass

from sanicx.router import params

# [stub]
sequence_shapes = {
    SHAPE_LIST,
    SHAPE_SET,
    SHAPE_TUPLE,
    SHAPE_SEQUENCE,
    SHAPE_TUPLE_ELLIPSIS,
}
sequence_types = (list, set, tuple)
sequence_shape_to_type = {
    SHAPE_LIST: list,
    SHAPE_SET: set,
    SHAPE_TUPLE: tuple,
    SHAPE_SEQUENCE: list,
    SHAPE_TUPLE_ELLIPSIS: list,
}


# [helper]
class AnnotionUtils:
    @classmethod
    def get_typed_annotation(
        cls, param: inspect.Parameter, globalns: Dict[str, Any]
    ) -> Any:
        """获取param参数的annotation，对于字符串需要解析"""
        annotation = param.annotation
        if isinstance(annotation, str):
            annotation = ForwardRef(annotation)  # type: ignore
            annotation = evaluate_forwardref(annotation, globalns, globalns)
        return annotation

    @classmethod
    def get_typed_signature(cls, call: Callable) -> inspect.Signature:
        """获取call的参数签名，对于使用字符串表示的对象需要解析成相应的类型"""
        signature = inspect.signature(call)
        globalns = getattr(call, "__globals__", {})
        typed_params = [
            inspect.Parameter(
                name=param.name,
                kind=param.kind,
                default=param.default,
                annotation=cls.get_typed_annotation(param, globalns),
            )
            for param in signature.parameters.values()
        ]
        typed_signature = inspect.Signature(typed_params)
        return typed_signature


class ModelTypeUtils:
    @classmethod
    def _create_field(
        cls,
        name: str,
        type_: Type[Any],
        class_validators: Optional[Dict[str, Validator]] = None,
        default: Optional[Any] = None,
        required: Union[bool, UndefinedType] = False,
        model_config: Type[BaseConfig] = BaseConfig,
        field_info: Optional[FieldInfo] = None,
        alias: Optional[str] = None,
    ) -> ModelField:
        """验证类型是否正确"""
        class_validators = class_validators or {}
        field_info = field_info or FieldInfo(None)

        response_field = functools.partial(
            ModelField,
            name=name,
            type_=type_,
            class_validators=class_validators,
            default=default,
            required=required,
            model_config=model_config,
            alias=alias,
        )

        try:
            return response_field(field_info=field_info)
        except RuntimeError:
            raise Exception(
                f"Invalid args for response field! Hint: check that {type_} is a valid pydantic field type"
            )

    @classmethod
    def is_scalar_field(cls, field: ModelField) -> bool:
        """
        判断是否是query, cookie, header
        """
        field_info = field.field_info
        if not (
            field.shape == SHAPE_SINGLETON
            and not lenient_issubclass(field.type_, BaseModel)
            and not lenient_issubclass(field.type_, sequence_types + (dict,))
            and not isinstance(field_info, (params.Body, params.FlaskParam))
        ):
            return False
        if field.sub_fields:
            if not all(cls.is_scalar_field(f) for f in field.sub_fields):
                return False
        return True

    @classmethod
    def is_scalar_sequence_field(cls, field: ModelField) -> bool:
        if (field.shape in sequence_shapes) and not lenient_issubclass(
            field.type_, BaseModel
        ):
            if field.sub_fields is not None:
                for sub_field in field.sub_fields:
                    if not cls.is_scalar_field(sub_field):
                        return False
            return True
        if lenient_issubclass(field.type_, sequence_types):
            return True
        return False

    @classmethod
    def get_param_field(
        cls,
        *,
        param: inspect.Parameter,
        param_name: str,
        default_field_info: Type[params.Param] = params.Param,
        force_type: Optional[params.ParamTypes] = None,
        ignore_default: bool = False,
    ) -> ModelField:
        """获取参数所对应的modelfield类型"""
        had_schema = False
        if not param.default == param.empty and ignore_default is False:
            default_value = param.default
        else:
            default_value = Required

        if isinstance(default_value, FieldInfo):
            # field_info
            had_schema = True
            field_info, default_value = default_value, default_value.default
            if (
                isinstance(field_info, params.Param)
                and getattr(field_info, "in_", None) is None
            ):
                field_info.in_ = default_field_info.in_
            if force_type:
                field_info.in_ = force_type  # type: ignore
        else:
            field_info = default_field_info(default_value)

        required = default_value == Required
        annotation: Any = Any if param.annotation == param.empty else param.annotation
        annotation = get_annotation_from_field_info(annotation, field_info, param_name)

        # 构造basemodel用于验证
        if fn := getattr(field_info, "validator", None):
            error_message = getattr(field_info, "error_message", "")
            annotation = type(
                "model",
                (BaseModel,),
                {
                    "__annotations__": {param_name: annotation},
                    "validate_name": ModelTypeUtils.build_validator_model(
                        param_name, fn, error_message
                    ),
                },
            )
        # 构造basemodel用于验证
        if not field_info.alias and getattr(field_info, "convert_underscores", None):
            alias = param.name.replace("_", "-")
        else:
            alias = field_info.alias or param.name
        field = cls._create_field(
            name=param.name,
            type_=annotation,
            default=None if required else default_value,
            alias=alias,
            required=required,
            field_info=field_info,
        )
        field.required = required
        if not had_schema and not cls.is_scalar_field(field=field):
            # 默认使用flaskparam
            field.field_info = params.FlaskParam(field_info.default)

        return field

    @classmethod
    def build_validator_model(cls, name: str, fn: Callable, error_message: str):
        @validator(name, allow_reuse=True)
        def _valid(cls, field: Any):
            message = error_message or f"[{name}]的值为[{field}]出错了"
            if not fn(field):
                raise ValueError(message)
            return field

        return _valid


class Parameter:
    """
    函数的参数类型
    """

    def __init__(self, container=None):
        self.inner_params: Optional[List[ModelField]] = []  # 由flask内置机制传入的类型, 包括路由参数
        self.query_params: Optional[List[ModelField]] = []  # 查询参数
        self.header_params: Optional[List[ModelField]] = []  # header请求头
        self.cookie_params: Optional[List[ModelField]] = []  # cookie参数
        self.body_params: Optional[List[ModelField]] = []  # json
        self.form_params: Optional[List[ModelField]] = []  # form表单数据
        self.file_params: Optional[List[ModelField]] = []  # file对象
        self.inject_params: Optional[List[inspect.Parameter]] = []  # 需要自动注入的对象

        self.background_tasks_param_name = None  # backgroundtask

    @classmethod
    def _add_param_to_fields(cls, *, field: ModelField, parameter: "Parameter") -> None:
        field_info = cast(params.Param, field.field_info)
        # if field_info.in_ == params.ParamTypes.path:
        #     parameter.path_params.append(field)
        if field_info.in_ == params.ParamTypes.query:
            parameter.query_params.append(field)
        elif field_info.in_ == params.ParamTypes.header:
            parameter.header_params.append(field)
        else:
            assert (
                field_info.in_ == params.ParamTypes.cookie
            ), f"non-body parameters must be in path, query, header or cookie: {field.name}"
            parameter.cookie_params.append(field)

    @classmethod
    def _validate_body(
        cls,
        required_params: List[ModelField],
        received_body: Optional[Mapping[str, Any]],
    ) -> Tuple[Dict[str, Any], List[ErrorWrapper]]:
        if not required_params:
            return {}, []
        values = {}
        errors = []

        field = required_params[0]
        field_info = field.field_info
        embed = getattr(field_info, "embed", None)
        field_alias_omitted = len(required_params) == 1 and not embed
        if field_alias_omitted:
            received_body = {field.alias: received_body}

        for field in required_params:
            loc: Tuple[str, ...]
            if field_alias_omitted:
                loc = ("body",)
            else:
                loc = ("body", field.alias)
            value: Optional[Any] = None

            if received_body is not None:
                try:
                    value = received_body.get(field.alias)
                except AttributeError:
                    errors.append(ErrorWrapper(MissingError(), loc=loc))
                    continue
            if (
                value is None
                or (isinstance(field_info, params.Form) and value == "")
                or (
                    isinstance(field_info, params.Form)
                    and field.shape in sequence_shapes
                    and len(value) == 0
                )
            ):
                if field.required:
                    errors.append(ErrorWrapper(MissingError(), loc=loc))
                else:
                    values[field.name] = deepcopy(field.default)
                continue

            v_, errors_ = field.validate(value, values, loc=loc)

            if isinstance(errors_, ErrorWrapper):
                errors.append(errors_)
            elif isinstance(errors_, list):
                errors.extend(errors_)
            else:
                values[field.name] = v_

        return values, errors

    @classmethod
    def _validate_param(
        cls, required_params: Sequence[ModelField], received_params: Mapping[str, Any]
    ) -> Tuple[Dict[str, Any], List[ErrorWrapper]]:
        values = {}
        errors = []

        for field in required_params:
            value = received_params.get(field.alias)
            field_info = field.field_info
            assert isinstance(
                field_info, params.Param
            ), "Params must be subclasses of Param"
            if value is None:
                if field.required:
                    errors.append(
                        ErrorWrapper(
                            MissingError(), loc=(field_info.in_.value, field.alias)
                        )
                    )
                else:
                    values[field.name] = deepcopy(field.default)
                continue
            v_, errors_ = field.validate(
                value, values, loc=(field_info.in_.value, field.alias)
            )
            if isinstance(errors_, ErrorWrapper):
                errors.append(errors_)
            elif isinstance(errors_, list):
                errors.extend(errors_)
            else:
                values[field.name] = v_
        return values, errors

    @classmethod
    def _validate_form(
        cls,
        required_params: Sequence[ModelField],
        received_files: Mapping[str, Any],
        type_name: str = "form",
    ) -> Tuple[Dict[str, Any], List[ErrorWrapper]]:
        values, errors = {}, []

        for field in required_params:
            loc = (type_name, field.alias)

            is_validator = getattr(field.field_info, "validator", None)
            value = None
            try:
                value = received_files.get(field.alias)
            except Exception:
                errors.append(ErrorWrapper(MissingError(), loc=loc))
                continue
            if value is None:
                errors.append(ErrorWrapper(MissingError(), loc=loc))
                continue

            if is_validator:
                v_, errors_ = field.validate({field.alias: value}, {}, loc=loc)
                values[field.name] = getattr(v_, field.alias, None)
                if errors_ is not None:
                    errors.extend(errors_.exc.raw_errors)
            else:
                v_, errors_ = field.validate(value, {}, loc=loc)
                values[field.name] = v_

        return values, errors

    @classmethod
    def _add_non_field_param_to_dependency(
        cls, *, param: inspect.Parameter, parameter: "Parameter"
    ) -> Optional[bool]:
        """一些额外的参数， 包括backgroundtasks"""
        if param.name == "request":  # request由sanic自己传入
            return True
        if isinstance(param.default, (Provide, Provider)):
            parameter.inject_params.append(param)
            return True
        return None

    @classmethod
    def factory(cls, fn: Callable, container=None):
        instance = Parameter(container)

        endpoint_signature = AnnotionUtils.get_typed_signature(fn)
        signature_parameters = endpoint_signature.parameters
        for param_name, param in signature_parameters.items():
            if cls._add_non_field_param_to_dependency(param=param, parameter=instance):
                continue
            # 判断参数类型，然后加入不同的app_param, path_params, query_params等
            param_field = ModelTypeUtils.get_param_field(
                param=param, default_field_info=params.FlaskParam, param_name=param_name
            )

            if ModelTypeUtils.is_scalar_field(field=param_field):
                cls._add_param_to_fields(field=param_field, parameter=instance)
            elif isinstance(
                param.default, (params.Query, params.Header)
            ) and ModelTypeUtils.is_scalar_sequence_field(param_field):
                cls._add_param_to_fields(field=param_field, parameter=instance)
            elif isinstance(param_field.field_info, params.File):
                instance.file_params.append(param_field)
            elif isinstance(param_field.field_info, params.Form):
                instance.form_params.append(param_field)
            elif isinstance(param_field.field_info, params.Body):
                instance.body_params.append(param_field)
            else:
                assert isinstance(
                    param_field.field_info, params.FlaskParam
                ), f"Param: {param_field.name} can only be a Flask inner arg, using FlaskParam(...)"
                instance.inner_params.append(param_field)

        return instance

    def validate(
        self,
        *,
        inner_params: Mapping[str, Any],
        query_params: Mapping[str, Any],
        header_params: Mapping[str, Any],
        cookie_params: Mapping[str, Any],
        body_params: Mapping[str, Any],
        file_params: Mapping[str, Any],
        form_params: Mapping[str, Any],
    ) -> Tuple[Dict[str, Any], List[ErrorWrapper]]:
        values: Dict[str, Any] = {}
        errors: List[ErrorWrapper] = []

        path_values, path_errors = self._validate_param(self.inner_params, inner_params)
        query_values, query_errors = self._validate_param(
            self.query_params, query_params
        )
        header_values, header_errors = self._validate_param(
            self.header_params, header_params
        )
        cookie_values, cookie_errors = self._validate_param(
            self.cookie_params, cookie_params
        )
        body_values, body_errors = self._validate_body(self.body_params, body_params)
        file_values, file_errors = self._validate_form(
            self.file_params, file_params, type_name="files"
        )
        form_values, form_errors = self._validate_form(self.form_params, form_params)

        values.update(path_values)
        values.update(query_values)
        values.update(header_values)
        values.update(cookie_values)
        values.update(body_values)
        values.update(file_values)
        values.update(form_values)
        errors += (
            path_errors
            + query_errors
            + header_errors
            + cookie_errors
            + body_errors
            + file_errors
            + form_errors
        )

        return values, errors
