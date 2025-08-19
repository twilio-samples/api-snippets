package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoicePay{
			PaymentMethod:  "security-code",
			ValidCardTypes: "visa mastercard amex",
			InnerElements: []twiml.Element{
				&twiml.VoicePrompt{
					For_: "payment-card-number",
					InnerElements: []twiml.Element{
						&twiml.VoiceSay{
							Message: "Please enter your credit card number.",
						},
					},
				},
				&twiml.VoicePrompt{
					For_:                  "payment-card-number",
					ErrorType:             "timeout",
					RequireMatchingInputs: "true",
					InnerElements: []twiml.Element{
						&twiml.VoiceSay{
							Message: "You didn't enter your credit card number. Please enter your credit card number.",
						},
					},
				},
				&twiml.VoicePrompt{
					ErrorType: "invalid-card-number",
					For_:      "payment-card-number",
					InnerElements: []twiml.Element{
						&twiml.VoiceSay{
							Message: "You didn't enter your credit card number. Please enter your credit card number.",
						},
					},
				},
				&twiml.VoicePrompt{
					ErrorType: "input-card-type",
					For_:      "payment-card-number",
					InnerElements: []twiml.Element{
						&twiml.VoiceSay{
							Message: "The card number you entered isn't from one of our accepted credit card issuers. Please enter a Visa, MasterCard, or American Express credit card number.",
						},
					},
				},
				&twiml.VoicePrompt{
					For_: "expiration-date",
					InnerElements: []twiml.Element{
						&twiml.VoiceSay{
							Message: "Please enter your credit card's expiration date. Two digits for the month and two digits for the year.",
						},
					},
				},
				&twiml.VoicePrompt{
					ErrorType:             "timeout",
					For_:                  "expiration-date",
					RequireMatchingInputs: "true",
					InnerElements: []twiml.Element{
						&twiml.VoiceSay{
							Message: "Sorry. You didn't enter an expiration date. Please enter your card's expiration date. Two digits for the month and two digits for the year.",
						},
					},
				},
				&twiml.VoicePrompt{
					ErrorType: "invalid-date",
					For_:      "expiration-date",
					InnerElements: []twiml.Element{
						&twiml.VoiceSay{
							Message: "The date you entered was incorrect or is in the past. Please enter the expiration date. Two digits for the month and two digits for the year. For example, to enter July twenty twenty two, enter 0 7 2 2.",
						},
					},
				},
				&twiml.VoicePrompt{
					CardType: "visa mastercard",
					For_:     "security-code",
					InnerElements: []twiml.Element{
						&twiml.VoiceSay{
							Message: "Please enter your security code. It's the 3 digits located on the back of your card.",
						},
					},
				},
				&twiml.VoicePrompt{
					CardType:  "visa mastercard",
					ErrorType: "timeout",
					For_:      "security-code",
					InnerElements: []twiml.Element{
						&twiml.VoiceSay{
							Message: "You didn't enter your credit card security code. Please enter your security code. It's the 3 digits located on the back of your card.",
						},
					},
				},
				&twiml.VoicePrompt{
					CardType:              "visa mastercard",
					ErrorType:             "invalid-security-code",
					For_:                  "security-code",
					RequireMatchingInputs: "true",
					InnerElements: []twiml.Element{
						&twiml.VoiceSay{
							Message: "That was an invalid security code. The security code must be 3 digits. Please try again.",
						},
					},
				},
				&twiml.VoicePrompt{
					CardType: "amex",
					For_:     "security-code",
					InnerElements: []twiml.Element{
						&twiml.VoiceSay{
							Message: "Please enter your security code. It's the 4 digits located on the front of your card.",
						},
					},
				},
				&twiml.VoicePrompt{
					CardType:  "amex",
					ErrorType: "timeout",
					For_:      "security-code",
					InnerElements: []twiml.Element{
						&twiml.VoiceSay{
							Message: "You didn't enter your credit card security code.  Please enter your security code. It's the 4 digits located on the front of your card.",
						},
					},
				},
				&twiml.VoicePrompt{
					CardType:  "amex",
					ErrorType: "invalid-security-code",
					For_:      "security-code",
					InnerElements: []twiml.Element{
						&twiml.VoiceSay{
							Message: "That was an invalid security code. The security code must be 4 digits. Please try again.",
						},
					},
				},
				&twiml.VoicePrompt{
					For_: "postal-code",
					InnerElements: []twiml.Element{
						&twiml.VoiceSay{
							Message: "Please enter your 5 digit billing zip code.",
						},
					},
				},
				&twiml.VoicePrompt{
					For_:      "postal-code",
					ErrorType: "timeout",
					InnerElements: []twiml.Element{
						&twiml.VoiceSay{
							Message: "You didn't enter your billing zip code. Please enter your 5 digit billing zip code.",
						},
					},
				},

				&twiml.VoicePrompt{
					For_: "payment-processing",
					InnerElements: []twiml.Element{
						&twiml.VoiceSay{
							Message: "Thank you. Please wait while we process your payment.",
						},
					},
				},
			},
		},
	})

	fmt.Print(twiml)
}
