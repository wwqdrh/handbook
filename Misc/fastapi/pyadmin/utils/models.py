from tortoise.models import Model as Model_
from tortoise.fields.data import IntField


class Model(Model_):
    id = IntField(pk=True)
