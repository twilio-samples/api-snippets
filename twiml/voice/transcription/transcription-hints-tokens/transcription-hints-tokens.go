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
					Hints:             "$OOV_CLASS_ALPHANUMERIC_SEQUENCE",
					StatusCallbackUrl: "https://example.com/your-callback-url",
				},
			},
		},
	})

	fmt.Print(twiml)
}
