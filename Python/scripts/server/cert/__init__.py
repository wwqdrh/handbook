import pathlib

_cur = pathlib.Path(__file__).parents[0]


def path(name: str) -> pathlib.Path:
    return _cur / name
