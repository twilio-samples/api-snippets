package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceStop{
			InnerElements: []twiml.Element{
				&twiml.VoiceTranscription{
					Name: "Contact center transcription",
				},
			},
		},
	})

	fmt.Print(twiml)
}
