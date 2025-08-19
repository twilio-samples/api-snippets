package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoicePay{
			ChargeAmount:  "13.22",
			PaymentMethod: "ach-debit",
			InnerElements: []twiml.Element{
				&twiml.VoicePrompt{
					For_: "bank-account-number",
					InnerElements: []twiml.Element{
						&twiml.VoiceSay{
							Message: "Thanks for using our service. Please enter your bank account number.",
						},
					},
				},
				&twiml.VoicePrompt{
					For_:                  "expiration-date",
					RequireMatchingInputs: "true",
					InnerElements: []twiml.Element{
						&twiml.VoiceSay{
							Message: "Thank you. Please enter your bank account number again.",
						},
					},
				},
			},
		},
	})

	fmt.Print(twiml)
}
