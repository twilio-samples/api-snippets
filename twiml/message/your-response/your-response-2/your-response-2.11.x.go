package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Messages([]twiml.Element{
		&twiml.MessagingMessage{
			Body: "This is message 1 of 2.",
		},
		&twiml.MessagingMessage{
			Body: "This is message 2 of 2.",
		},
	})

	fmt.Print(twiml)
}
