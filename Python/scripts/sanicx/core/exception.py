class UseCaseException(Exception):
    def __init__(self, message, code: int = 50000):
        super().__init__(message)
        self.message = message
        self.code = code


class FileExtensionException(UseCaseException):
    def __init__(self, message: str):
        super().__init__(message, 50021)


class ImproperlyConfigured(Exception):
    """Django is somehow improperly configured"""
    pass


class CommandError(Exception):
    """
    解析命令行参数出错
    """
    pass
