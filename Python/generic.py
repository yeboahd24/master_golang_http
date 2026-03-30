from typing import TypeVar

Number = TypeVar("Number", int, float)


def add(a: Number, b: Number) -> Number:
    return a + b


if __name__ == "__main__":
    print(add(1, 2))
    print(add(1.0, 2.0))
    print(add(1, 2.0))  # error but it will be converted to float
