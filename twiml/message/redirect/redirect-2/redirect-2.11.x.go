package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Messages([]twiml.Element{
		&twiml.MessagingRedirect{
			Url: "../nextInstructions",
		},
	})

	fmt.Print(twiml)
}
