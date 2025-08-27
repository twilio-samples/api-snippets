package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Messages([]twiml.Element{
		&twiml.MessagingMessage{
			Body: "Store Location: 123 Easy St.",
		},
	})

	fmt.Print(twiml)
}
