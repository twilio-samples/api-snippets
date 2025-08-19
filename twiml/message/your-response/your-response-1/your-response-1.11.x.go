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
					Message: "Hello World!",
				},
			},
		},
		&twiml.MessagingRedirect{
			Url: "https://demo.twilio.com/welcome/sms/",
		},
	})

	fmt.Print(twiml)
}
