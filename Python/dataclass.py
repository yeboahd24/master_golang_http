from dataclasses import dataclass, field, InitVar, replace
from typing import ClassVar
from functools import cached_property
from datetime import datetime


@dataclass(frozen=False, slots=True)
class User:
    id: int
    username: str
    email: str
    created_at: datetime = field(default_factory=datetime.now)

    _hashed_password: int | None = field(init=False, repr=False)

    # Class variable (shared between instances)
    MAX_LOGIN_ATTEMPTS: ClassVar[int] = 5

    # Only used during __init__, not stored
    password_raw: InitVar[str | None] = None

    @cached_property
    def display_name(self) -> str:
        return f"{self.username} <{self.email}>"

    # We did this to mutate because the frozen=True won't allow setting attr
    # def __post_init__(self, password_raw):
    #     object.__setattr__(
    #         self, "_hashed_password", hash(password_raw) if password_raw else None
    #     )

    # This will cause an error because of the fronzen=True unless set to false which is default behavior
    def __post_init__(self, password_raw: str | None):
        if password_raw:
            self._hashed_password = hash(password_raw)
        else:
            self._hashed_password = None


if __name__ == "__main__":
    admin = User(1, "admin", "adminc@gamil.com")
    updated = replace(admin, username="Dominic")
    print(admin)
    print(updated)
