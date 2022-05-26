from logging import handlers
import pathlib
import datetime
import logging
import os

fmt = logging.Formatter("%(asctime)s %(levelname)s: %(message)s", "%Y-%m-%d %H:%M:%S")


def init_logger(app: str, dirs: pathlib.Path) -> logging.Logger:
    """
    the app logger
    """
    os.makedirs(str(dirs), mode=0o740, exist_ok=True)
    appLogger = logging.getLogger(app)
    appLogger.setLevel(logging.INFO)  # info以上的才会输出

    # info level to console
    console = logging.StreamHandler()
    console.setFormatter(fmt)
    console.setLevel(logging.INFO)

    # error level to file
    today = datetime.date.today()
    file_name = dirs / ("exceptions_" + str(today) + ".log")
    fh = handlers.TimedRotatingFileHandler(
        filename=file_name, when="D", backupCount=30, encoding="utf-8"
    )
    fh.setFormatter(fmt)
    fh.setLevel(logging.ERROR)

    # add handler
    appLogger.addHandler(console)
    appLogger.addHandler(fh)

    return appLogger
