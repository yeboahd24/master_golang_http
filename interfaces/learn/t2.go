package main

import "fmt"

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

// Without an interface, you're forced to know about every single type
func AlertAll(emails []EmailNotifier, slacks []SlackNotifier, smses []SMSNotifier, message string) {
	for _, e := range emails {
		e.Send(message)
	}
	for _, s := range slacks {
		s.Send(message)
	}
	for _, s := range smses {
		s.Send(message)
	}
}

func main() {
	emails := []EmailNotifier{{Address: "dev@company.com"}}
	slacks := []SlackNotifier{{Channel: "alerts"}}
	smses := []SMSNotifier{{Phone: "+1-555-1234"}}

	AlertAll(emails, slacks, smses, "Server is on fire!")
}
