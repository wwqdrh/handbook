"""
基于sqlalchemy的orm框架
"""
import contextlib
import typing as T
from collections import UserDict
from contextlib import asynccontextmanager

import alembic
import sqlalchemy
from asyncref import adapter
from sqlalchemy import Table, Column
from sqlalchemy.engine import Engine  # type: ignore
from sqlalchemy.ext.declarative import DeclarativeMeta, declarative_base
from sqlalchemy.orm import Session, sessionmaker, Query

__all__ = ("Repo",)

metadata = sqlalchemy.MetaData()


class RepoMeta(type):
    _table: sqlalchemy.Table
    _fields: T.List[str]
    _migrate: T.Callable

    class member(UserDict):
        def __init__(self):
            super().__init__()
            self.fields = []
            self.table_name = None

        def __setitem__(self, key, value):
            if isinstance(value, Column):
                self.fields.append(key)
            elif key == "__tablename__":
                self.table_name = value

            super().__setitem__(key, value)

    @classmethod
    def __prepare__(mcs, cls_name: str, bases: tuple) -> member:
        return mcs.member()

    def __new__(mcs, cls_name: str, bases: tuple, class_dict: member):
        result = type.__new__(mcs, cls_name, bases, dict(class_dict))
        result._table = sqlalchemy.Table(
            class_dict.table_name,
            metadata,
            *map(lambda i: class_dict[i], class_dict.fields),
        )
        result._fields = class_dict.fields
        result._migrate = lambda: alembic.op.create_module(
            class_dict.table_name, *map(lambda i: class_dict[i], class_dict.fields)
        )
        return result

    def __call__(cls, engine: Engine) -> "RepoMeta":
        # 实例化内容
        ins = cls.__new__(cls)
        ins.__init__(cls._table, cls._migrate, engine)
        return ins


class Repo(metaclass=RepoMeta):
    def __init__(self, table: Table, migrate, engine: Engine):
        self.engine = engine
        self.session = sessionmaker(bind=engine)
        self.table = table

        self._migrate = migrate

    @asynccontextmanager
    async def transactions(self) -> T.AsyncGenerator[Session, None]:
        session: Session = self.session()
        try:
            yield session
        except Exception as e:
            session.rollback()
            raise e
        else:
            session.commit()
        finally:
            session.close()

    async def execute(self, sm, session: Session = None):
        if session is None:
            with self.engine.connect() as conn:
                result = await adapter.wrap(conn.execute, sm)
        else:
            result = await adapter.wrap(session.execute, sm)
        return result

    async def create(self, session: Session = None, **data) -> int:
        for field in self._fields:
            if act := getattr(self, f"gen_{field}", None):
                val = act()
                if val is not None:
                    data[field] = val

        result = await self.execute(
            sqlalchemy.sql.insert(self.table).values(**data), session
        )
        return result.lastrowid

    async def delete(self, column_value, column_name="id", session: Session = None):
        result = await self.execute(
            sqlalchemy.sql.delete(self.table).where(
                self.table.c[column_name] == column_value
            ),
            session,
        )

        return result.rowcount

    async def modify(
        self, column_value, column_name="id", session: Session = None, **data
    ):
        data = {k: v for k, v in data.items() if v is not None}

        await self.execute(
            sqlalchemy.sql.update(self.table)
            .where(self.table.c[column_name] == column_value)
            .values(**data),
            session,
        )

        return await self.info(column_value, column_name)

    async def info(self, column_value, column_name="id", session: Session = None):
        if column_value is None:
            return None
        with contextlib.suppress(BaseException):
            result = await self.execute(
                self.table.select().where(self.table.c[column_name] == column_value),
                session,
            )
            row = result.first()

            return None if row is None else dict(row)
        return None

    async def infos(self, column_values, column_name="id"):
        valid_values = [v for v in column_values if v is not None]
        if valid_values:
            result = await self.execute(
                self.table.select().where(self.table.c[column_name].in_(valid_values))
            )
            d = {v[column_name]: dict(v) for v in result.fetchall()}
        else:
            d = {}

        return [d.get(v) for v in column_values]

    async def list(
        self, *, from_=None, where=None, order_by=None, limit=None, offset=None
    ):
        select_sm = self.table.select()
        count_sm = sqlalchemy.sql.select([sqlalchemy.sql.func.count()]).select_from(
            self.table
        )

        if from_ is not None:
            select_sm = select_sm.select_from(from_)
            count_sm = count_sm.select_from(from_)

        if where is not None:
            select_sm = select_sm.where(where)
            count_sm = count_sm.where(where)

        if order_by is not None:
            select_sm = select_sm.order_by(order_by)

        if limit is not None:
            select_sm = select_sm.limit(limit)
        if offset is not None:
            select_sm = select_sm.offset(offset)

        result = await self.execute(select_sm)
        rows = [dict(v) for v in result.fetchall()]

        result = await self.execute(count_sm)
        total = result.scalar()

        return rows, total

    async def count(self, where=None):
        sm = sqlalchemy.sql.select([sqlalchemy.sql.func.count()]).select_from(
            self.table
        )
        if where is not None:
            sm = sm.where(where)
        result = await self.execute(sm)

        return result.scalar()
