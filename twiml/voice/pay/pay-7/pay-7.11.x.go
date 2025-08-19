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
					CardType: "amex",
					For_:     "security-code",
					InnerElements: []twiml.Element{
						&twiml.VoiceSay{
							Message: "Please enter security code for your American Express card. It’s the 4 digits located on the front of your card",
						},
					},
				},
			},
		},
	})

	fmt.Print(twiml)
}
