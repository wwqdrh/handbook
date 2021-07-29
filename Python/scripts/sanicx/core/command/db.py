import argparse
import typing as T

from sanicx.globals import current_app

if T.TYPE_CHECKING:
    from sanicx.db.migrator import MigrateConfig

CORE_COMMAND = "DBCommand"


class DBCommand:
    command = argparse.ArgumentParser(add_help=False)

    command.add_argument('command', nargs="?", choices=["init", "upgrade", "downgrade"])

    # init
    init = command.add_argument_group("init", "初始化migrate")
    init.add_argument('-d', '--directory', action="store_true",
                      help='Migration script directory (default is "migrations")')
    init.add_argument('--multidb', action="store_true", help="'Support multiple databases'")

    # upgrade
    upgrade = command.add_argument_group("upgrade|downgrade", "升级|降低版本")
    upgrade.add_argument('--sql', action="store_true",
                         help=('Don\'t emit SQL to database - dump to standard output '
                               'instead'))
    upgrade.add_argument('--tag', action="store_true",
                         help=('Arbitrary "tag" name - can be used by custom env.py '
                               'scripts'))
    upgrade.add_argument('revision', nargs='?', default="-1", help="revision identifier")
    upgrade.add_argument('-x', '--x-arg', dest='x_arg', default=None,
                         action='append', help=("Additional arguments consumed "
                                                "by custom env.py scripts"))

    command.add_argument('--autogenerate', action="store_true",
                         help=('Populate revision script with candidate migration '
                               'operations, based on comparison of database to model'))
    command.add_argument('--head', default='head',
                         help=('Specify head revision or <branchname>@head to base new '
                               'revision on'))
    command.add_argument('--splice', action="store_true",
                         help='Allow a non-head revision as the "head" to splice onto')
    command.add_argument('--branch-label', action="store_true",
                         help='Specify a branch label to apply to the new revision')
    command.add_argument('--version-path', action="store_true",
                         help='Specify specific path from config for version file')
    command.add_argument('--rev-id', action="store_true",
                         help=('Specify a hardcoded revision id instead of generating '
                               'one'))
    command.add_argument('-r', '--rev-range', default=None,
                         help='Specify a revision range; format is [start]:[end]')
    command.add_argument('-v', '--verbose', action="store_true", help='Use more verbose output')
    command.add_argument('-i', '--indicate-current', action="store_true",
                         help=('Indicate current version (Alembic 0.9.9 or greater is '
                               'required)'))
    command.add_argument('--resolve-dependencies', action="store_true",
                         help='Treat dependency versions as down revisions')
    command.add_argument('--head-only', action="store_true",
                         help='Deprecated. Use --verbose for additional output')

    def __init__(self, args: T.List[str]):
        self.parser: argparse.Namespace = self.command.parse_args(args)
        self.migrate: "MigrateConfig" = current_app.extensions["migrate"]

    def __call__(self):
        parser, command = self.parser, self.parser.command
        print(command)
        if command == "init":
            return self.migrate.init(parser.directory, parser.multidb)
        elif command == "upgrade":
            return self.migrate.upgrade(parser.tag, parser.sql, parser.revision, parser.directory, parser.x_arg)
        elif command == "downgrade":
            return self.migrate.downgrade(parser.tag, parser.sql, parser.revision, parser.directory, parser.x_arg)
