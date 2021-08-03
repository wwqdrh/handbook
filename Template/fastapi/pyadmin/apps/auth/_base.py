from fastapi import HTTPException, status

credentials_exception = lambda auth_value: HTTPException(
    status_code=status.HTTP_401_UNAUTHORIZED,
    detail="Could not validate credentials",
    headers={"WWW-Authenticate": auth_value},
)

no_permission_exception = lambda auth_value: HTTPException(
    status_code=status.HTTP_401_UNAUTHORIZED,
    detail="Not enough permissions",
    headers={"WWW-Authenticate": auth_value},
)