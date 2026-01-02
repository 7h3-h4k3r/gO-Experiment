package main

import (
	"fmt"
)

type Notify interface {
	Notifying(message string) error
}

type EmailNotify struct {
	Email string
}

type PhoneNotify struct {
	phone string
	code  int
}

func (e EmailNotify) Notifying(message string) error {
	fmt.Printf("Email %s send to message %s ", e.Email, message)
	return nil
}

func (e PhoneNotify) Notifying(message string) error {
	fmt.Printf("Message SMS send the phone number %s [code]:%d", e.phone, e.code)
	return nil
}
func sendNotify(n Notify, msg string) {
	err := n.Notifying(msg)
	if err != nil {
		fmt.Println("Error :", err)
	}
}

func main() {
	notifier := EmailNotify{Email: "dharani@example.com"}
	sms := PhoneNotify{phone: "+9080443868", code: 345236}
	sendNotify(notifier, "hi dharani what about the go project how can you interact with them")
	sendNotify(sms, "hi dharani what about the go project how can you interact with them")
}
