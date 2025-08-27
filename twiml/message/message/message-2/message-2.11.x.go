package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Messages([]twiml.Element{
		&twiml.MessagingMessage{
			InnerElements: []twiml.Element{
				&twiml.MessagingBody{
					Message: "Store Location: 123 Easy St.",
				},
				&twiml.MessagingMedia{
					Url: "https://demo.twilio.com/owl.png",
				},
			},
		},
	})

	fmt.Print(twiml)
}
