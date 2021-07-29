from pyadmin.context import current_app, Application


def app() -> Application:
    assert (app := current_app.get(None)) is not None, "未正常启动应用"

    return app
