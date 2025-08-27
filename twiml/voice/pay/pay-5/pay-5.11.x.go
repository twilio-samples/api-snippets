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
						&twiml.VoicePlay{
							Url: "https://myurl.com/twilio/twiml/audio/card_number.mp3",
						},
					},
				},
			},
		},
	})

	fmt.Print(twiml)
}
