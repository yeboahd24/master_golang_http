# In duck typing,  we don’t care about the object’s type;
# we care about whether it can do what we need it to do.


from typing import Protocol


class Notifier(Protocol):
    def send_notification(self, message: str) -> None: ...


class EmailNotifier:  # Note: no explicit inheritance
    def send_notification(self, message: str) -> None:
        print(f"Sending email: {message}")


class SMSNotifier:  # Note: no explicit inheritance
    def send_notification(self, message: str) -> None:
        print(f"Sending SMS: {message}")


class NotificationService:
    # Still able to use type hinting
    def __init__(self, notifier: Notifier):
        self.notifier = notifier

    def notify(self, message: str) -> None:
        self.notifier.send_notification(message)


if __name__ == "__main__":
    email_notifier = EmailNotifier()
    sms_notifier = SMSNotifier()
    notification_service = NotificationService(email_notifier)
    notification_service.notify("Hello, world!")
