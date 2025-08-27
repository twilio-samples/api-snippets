package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Messages([]twiml.Element{
		&twiml.MessagingMessage{
			Action: "/SmsHandler.php",
			Body:   "Store Location: 123 Easy St.",
			Method: "POST",
		},
	})

	fmt.Print(twiml)
}
