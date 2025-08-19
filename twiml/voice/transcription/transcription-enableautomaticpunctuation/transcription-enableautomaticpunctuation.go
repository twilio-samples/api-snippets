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
					EnableAutomaticPunctuation: "true",
					StatusCallbackUrl:          "https://example.com/your-callback-url",
					TranscriptionEngine:        "google",
				},
			},
		},
	})

	fmt.Print(twiml)
}
