# Payment Processor Interface
from abc import ABC, abstractmethod
from decimal import Decimal


class PaymentProcessor(ABC):
    @abstractmethod
    def process(self, amount: Decimal) -> bool:
        pass

    @abstractmethod
    def refund(self, amount: Decimal) -> bool:
        pass


class MobileMoneyProcessor(PaymentProcessor):
    def process(self, amount: Decimal) -> bool:
        print(f"Processing MOMO payment of {amount}")
        return True

    def refund(self, amount: Decimal) -> bool:
        print(f"Refunding {amount} via Mobile Money")
        return True


if __name__ == "__main__":
    momo = MobileMoneyProcessor()
    momo.process(200)
    momo.refund(200)
