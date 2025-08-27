package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoicePay{
			InnerElements: []twiml.Element{
				&twiml.VoicePrompt{
					For_: "payment-card-number",
					InnerElements: []twiml.Element{
						&twiml.VoiceSay{
							Message: "Please enter security code for your Visa card. It’s the 3 digits located on the back of your card",
						},
					},
				},
			},
		},
	})

	fmt.Print(twiml)
}
