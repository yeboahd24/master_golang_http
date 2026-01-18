# Payment Processor Interface
from typing import Protocol


class PaymentProcessor(Protocol):
    def charge(self, amount: int) -> bool:
        pass

    def refund(self, amount: int) -> bool:
        pass


class MobileMoneyProcessor(PaymentProcessor):
    def charge(self, amount: int) -> bool:
        print(f"Charging MOMO payment of {amount}")
        return True

    def refund(self, amount: int) -> bool:
        print(f"Refunding {amount} via Mobile Money")
        return True


# Old Legacy Code
class OldLegacyPayPalAdapter:
    def charge(self, cent: int) -> bool:
        print(f"Charge PayPal payment of {cent}")
        return True

    def refund(self, cent: int) -> bool:
        print(f"Refunding {cent} via Mobile Money")
        return True


def process_payment(p: PaymentProcessor, amount: int):
    p.charge(amount)


if __name__ == "__main__":
    process_payment(MobileMoneyProcessor(), 500)
