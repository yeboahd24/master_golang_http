package main

// O — Open/Closed Principle
// Open for extension, closed for modification.
// In Go, you achieve this through interfaces — you extend behaviour by implementing a new type, not by editing existing ones.

type Notifier interface {
	Notify(msg string) error
}

// Add new notification channels without touching existing code
type EmailNotifier struct{}

func (e EmailNotifier) Notify(msg string) error { /* send email */ }

type SMSNotifier struct{}

func (s SMSNotifier) Notify(msg string) error { /* send SMS */ }

type SlackNotifier struct{}

func (s SlackNotifier) Notify(msg string) error { /* send Slack */ }

// This function never changes, regardless of how many Notifiers you add
func AlertOps(n Notifier, msg string) error {
	return n.Notify(msg)
}
