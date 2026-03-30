from typing import NewType, Literal

UserID = NewType("UserID", int)
ProductID = NewType("ProductID", int)

LogLevel = Literal["DEBUG", "INFO", "WARNING", "ERROR", "CRITICAL"]


def set_log_level(level: LogLevel) -> None:
    print(f"Setting log level to {level}")


def process_order(user_id: UserID, product_id: ProductID) -> None:
    print(f"Processing order for user {user_id} and product {product_id}")


if __name__ == "__main__":
    process_order(UserID(1), ProductID(2))
    process_order(ProductID(1), UserID(2))  # This will raise a TypeError

    print("-----------------------------------------------")
    set_log_level("DEBUG")
    set_log_level("WRONG")  # This will raise a ValueError
