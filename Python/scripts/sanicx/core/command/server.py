import argparse

from sanicx import globals

CORE_COMMAND = "command"


def command(args: argparse.Namespace):
    app = globals.current_app
    app.run(host="0.0.0.0", port=8080)
