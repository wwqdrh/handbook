from contextvars import ContextVar

from pyadmin.main import Application

current_app: ContextVar[Application] = ContextVar("current_app")
