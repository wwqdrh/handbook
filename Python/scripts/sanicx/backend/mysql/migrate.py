"""
migrate操作，模型以及版本

from alembic import op
import sqlalchemy as sa

def upgrade():
    op.create_table(
        "account",
        sa.Column("id", sa.Integer, primary_key=True),
        sa.Column("name", sa.String(50), nullable=False),
        sa.Column("description", sa.Unicode(200)),
    )


def downgrade():
    op.drop_table("account")

"""
import argparse
import os

from alembic import __version__ as __alembic_version__
from alembic import command
from alembic.config import Config as AlembicConfig

__all__ = ("Migrate",)

alembic_version = tuple([int(v) for v in __alembic_version__.split('.')[0:3]])


class Config(AlembicConfig):
    def get_template_directory(self):
        package_dir = os.path.abspath(os.path.dirname(__file__))
        return os.path.join(package_dir, 'templates')


class MigrateConfig:
    """
    Methods:
        init: Create a new migration repository
        branches: show current branch points
        current: display the current revision for each database
        downgrade: Revert to a previous version
        edit: Edit a revision file
        heads: Show current available heads in the script...
        history: List changeset scripts in chronological order.
        merge: Merge two revisions together, creating a new...
        migrate: Autogenerate a new revision file (Alias for...
        revision: Create a new revision file.
        show: Show the revision denoted by the given symbol.
        stamp: 'stamp' the revision table with the given revision;...
        upgrade: Upgrade to a later version
    """

    def __init__(self, migrate, db, **kwargs):
        self.migrate = migrate
        self.db = db
        self.directory = migrate.directory
        self.configure_args = kwargs

    @property
    def metadata(self):
        """
        Backwards compatibility, in old releases app.extensions['migrate']
        was set to db, and env.py accessed app.extensions['migrate'].metadata
        """
        return self.db.metadata

    def init(self, directory=None, multidb=False):
        """Creates a new migration repository"""
        directory = self.directory
        config = Config()
        config.set_main_option('script_location', directory)
        config.config_file_name = os.path.join(directory, 'alembic.ini')
        config = self.migrate.call_configure_callbacks(config)
        if multidb:
            command.init(config, directory, 'flask-multidb')
        else:
            command.init(config, directory, 'flask')

    def revision(self, directory=None, message=None, autogenerate=False, sql=False,
                 head='head', splice=False, branch_label=None, version_path=None,
                 rev_id=None):
        """Create a new revision file."""
        config = self.migrate.get_config(directory)
        command.revision(config, message, autogenerate=autogenerate, sql=sql,
                         head=head, splice=splice, branch_label=branch_label,
                         version_path=version_path, rev_id=rev_id)

    def migrate(self, directory=None, message=None, sql=False, head='head', splice=False,
                branch_label=None, version_path=None, rev_id=None, x_arg=None):
        """Alias for 'revision --autogenerate'"""
        config = self.migrate.get_config(
            directory, opts=['autogenerate'], x_arg=x_arg)
        command.revision(config, message, autogenerate=True, sql=sql,
                         head=head, splice=splice, branch_label=branch_label,
                         version_path=version_path, rev_id=rev_id)

    def edit(self, directory=None, revision='current'):
        """Edit current revision."""
        if alembic_version >= (0, 8, 0):
            config = self.migrate.get_config(
                directory)
            command.edit(config, revision)
        else:
            raise RuntimeError('Alembic 0.8.0 or greater is required')

    def upgrade(self, directory=None, revision='head', sql=False, tag=None, x_arg=None):
        """Upgrade to a later version"""
        config = self.migrate.get_config(directory,
                                         x_arg=x_arg)
        command.upgrade(config, revision, sql=sql, tag=tag)

    def downgrade(self, directory=None, revision='-1', sql=False, tag=None, x_arg=None):
        """Revert to a previous version"""
        config = self.migrate.get_config(directory,
                                         x_arg=x_arg)
        if sql and revision == '-1':
            revision = 'head:-1'
        command.downgrade(config, revision, sql=sql, tag=tag)

    def show(self, directory=None, revision='head'):
        """Show the revision denoted by the given symbol."""
        config = self.migrate.get_config(directory)
        command.show(config, revision)

    def history(self, directory=None, rev_range=None, verbose=False,
                indicate_current=False):
        """List changeset scripts in chronological order."""
        config = self.migrate.get_config(directory)
        if alembic_version >= (0, 9, 9):
            command.history(config, rev_range, verbose=verbose,
                            indicate_current=indicate_current)
        else:
            command.history(config, rev_range, verbose=verbose)

    def heads(self, directory=None, verbose=False, resolve_dependencies=False):
        """Show current available heads in the script directory"""
        config = self.migrate.get_config(directory)
        command.heads(config, verbose=verbose,
                      resolve_dependencies=resolve_dependencies)

    def branches(self, directory=None, verbose=False):
        """Show current branch points"""
        config = self.migrate.get_config(directory)
        command.branches(config, verbose=verbose)

    def current(self, directory=None, verbose=False, head_only=False):
        """Display the current revision for each database."""
        config = self.migrate.get_config(directory)
        command.current(config, verbose=verbose, head_only=head_only)

    def stamp(self, directory=None, revision='head', sql=False, tag=None):
        """'stamp' the revision table with the given revision; don't run any
        migrations"""
        config = self.migrate.get_config(directory)
        command.stamp(config, revision, sql=sql, tag=tag)


class Migrate:
    """
    插件类
    """

    def __init__(self, app=None, db=None, directory='migrations', **kwargs):
        self.configure_callbacks = []
        self.db = db
        self.directory = str(directory)
        self.alembic_ctx_kwargs = kwargs
        if app is not None and db is not None:
            self.init_app(app, db, directory)

    def init_app(self, app, db=None, directory=None, **kwargs):
        self.db = db or self.db
        self.directory = str(directory or self.directory)
        self.alembic_ctx_kwargs.update(kwargs)
        if not hasattr(app, 'extensions'):
            app.extensions = {}
        app.extensions['migrate'] = MigrateConfig(
            self, self.db, **self.alembic_ctx_kwargs)

    def configure(self, f):
        self.configure_callbacks.append(f)
        return f

    def call_configure_callbacks(self, config):
        for f in self.configure_callbacks:
            config = f(config)
        return config

    def get_config(self, directory=None, x_arg=None, opts=None):
        if directory is None:
            directory = self.directory
        directory = str(directory)
        config = Config(os.path.join(directory, 'alembic.ini'))
        config.set_main_option('script_location', directory)
        if config.cmd_opts is None:
            config.cmd_opts = argparse.Namespace()
        for opt in opts or []:
            setattr(config.cmd_opts, opt, True)
        if not hasattr(config.cmd_opts, 'x'):
            if x_arg is not None:
                setattr(config.cmd_opts, 'x', [])
                if isinstance(x_arg, list) or isinstance(x_arg, tuple):
                    for x in x_arg:
                        config.cmd_opts.x.append(x)
                else:
                    config.cmd_opts.x.append(x_arg)
            else:
                setattr(config.cmd_opts, 'x', None)
        return self.call_configure_callbacks(config)
