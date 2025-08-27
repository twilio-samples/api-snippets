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
					Attempt: "1",
					For_:    "expiration-date",
					InnerElements: []twiml.Element{
						&twiml.VoiceSay{
							Message: "Please enter your expiration date, two digits for the month and two digits for the year.",
						},
					},
				},
				&twiml.VoicePrompt{
					Attempt: "2 3",
					For_:    "expiration-date",
					InnerElements: []twiml.Element{
						&twiml.VoiceSay{
							Message: "Please enter your expiration date, two digits for the month and two digits for the year. For example, if your expiration date is March 2022, then please enter 0 3 2 2",
						},
					},
				},
			},
		},
	})

	fmt.Print(twiml)
}
