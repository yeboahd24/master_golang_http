package main

import "fmt"

// The interface — anything that can Send a message is a Notifier
type Notifier interface {
    Send(message string) error
}

// --- Three different implementations ---

type EmailNotifier struct {
    Address string
}

func (e EmailNotifier) Send(message string) error {
    fmt.Printf("[EMAIL to %s]: %s\n", e.Address, message)
    return nil
}

type SlackNotifier struct {
    Channel string
}

func (s SlackNotifier) Send(message string) error {
    fmt.Printf("[SLACK #%s]: %s\n", s.Channel, message)
    return nil
}

type SMSNotifier struct {
    Phone string
}

func (s SMSNotifier) Send(message string) error {
    fmt.Printf("[SMS to %s]: %s\n", s.Phone, message)
    return nil
}

// --- The power of the interface ---
// This function doesn't know or care HOW the notification is sent.
// It just knows the thing it received CAN send.
func AlertAll(notifiers []Notifier, message string) {
    for _, n := range notifiers {
        n.Send(message)
    }
}

func main() {
    notifiers := []Notifier{
        EmailNotifier{Address: "dev@company.com"},
        SlackNotifier{Channel: "alerts"},
        SMSNotifier{Phone: "+1-555-1234"},
    }

    AlertAll(notifiers, "Server is on fire!")
}
```

This prints:
```
[EMAIL to dev@company.com]: Server is on fire!
[SLACK #alerts]: Server is on fire!
[SMS to +1-555-1234]: Server is on fire!
