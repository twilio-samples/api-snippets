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
							Message: "Please enter your 16 digit Visa or Mastercard number.",
						},
					},
				},
			},
		},
	})

	fmt.Print(twiml)
}
