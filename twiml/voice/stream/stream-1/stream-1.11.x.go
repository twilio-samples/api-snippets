package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceStart{
			InnerElements: []twiml.Element{
				&twiml.VoiceStream{
					Name: "My SIPREC Stream",
					Url:  "wss://example.com/audiostream",
				},
			},
		},
		&twiml.VoiceSay{
			Message: "The stream has started.",
		},
	})

	fmt.Print(twiml)
}
