package main

import (
	"fmt"
	"time"
)

type Notifier interface {
	Notify(message string) error
}

type Logger struct {
	prefix string
}
type EmailNotify struct {
	Logger
	Email string
}

func (l Logger) Log(msg string) {
	fmt.Printf("%s [%s]: %s\n", l.prefix, time.Now().Format(time.RFC3339), msg)
}

func (e EmailNotify) Notify(messae string) error {
	e.Log(fmt.Sprintf("Sending email to %s : %s", e.Email, messae))
	return nil
}

func sendNotification(n Notifier, msg string) {
	if err := n.Notify(msg); err != nil {
		fmt.Printf("notification Failed", err)
	}
}

func main() {
	email := EmailNotify{Logger: Logger{prefix: "Email Address"}, Email: "dharani@gmail.com"}
	sendNotification(email, "hi da")
}
