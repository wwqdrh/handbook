import hashlib


def get_hash_password(password: str, salt: str) -> str:
    sec = hashlib.md5()
    sec.update(f"{password}{salt}".encode("utf8"))
    sec.update(sec.hexdigest().encode("utf8"))
    return sec.hexdigest()