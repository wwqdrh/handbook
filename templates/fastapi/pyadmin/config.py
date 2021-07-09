from starlette.config import Config
from starlette.datastructures import Secret


from pyadmin.utils.configs import DatabaseURL

ENV = Config(".env")

MODE = Config(".{}.env".format(ENV.get("mode", str, "dev")))


DB_URL = DatabaseURL(
    drivername=MODE.get("DB_DRIVERNAME", str, "")
    or ENV.get("DB_DRIVERNAME", str, "")
    or "mysql",
    host=MODE.get("DB_HOST", str, "") or ENV.get("DB_HOST", str, "") or "localhost",
    port=MODE.get("DB_PORT", int, 0) or ENV.get("DB_PORT", int, 0) or 3306,
    username=MODE.get("DB_USERNAME", str, "") or ENV.get("DB_USERNAME", str, "") or "",
    password=MODE.get("DB_PASSWORD", Secret, "")
    or ENV.get("DB_PASSWORD", Secret, "")
    or Secret(""),
    database=MODE.get("DB_DATABASE", str, "") or ENV.get("DB_DATABASE", str, "") or "",
)

MODELS = [
    "aerich.models",
    "pyadmin.models.repo.admin",
]

DB_ORM = {
    "connections": {"default": DB_URL.url},
    "apps": {"default": {"models": MODELS, "default_connection": "default"}},
}


# addr = "127.0.0.1:6379"
#   db = "0"
#   maxretries = 3
#   minidleconns = 5
#   pass = "123456"
#   poolsize = 10
REDIS = {"addr": "127.0.0.1:6379", "db": "0", "pass": "123456", "poolsize": 10}


SECRET_KEY = "09d25e094faa6ca2556c818166b7a9563b93f7099f6f0f4caa6cf63b88e8d3e7"  # openssl rand -hex 32
ALGORITHM = "HS256"
ACCESS_TOKEN_EXPIRE_MINUTES = 30

SALT = "c04d20ed6e56b62d"