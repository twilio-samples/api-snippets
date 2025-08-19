package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceStart{
			InnerElements: []twiml.Element{
				&twiml.VoiceTranscription{
					Hints:             "Alice Johnson, Bob Martin, ACME Corp, XYZ Enterprises, product demo, sales inquiry, customer feedback",
					StatusCallbackUrl: "https://example.com/your-callback-url",
				},
			},
		},
	})

	fmt.Print(twiml)
}
