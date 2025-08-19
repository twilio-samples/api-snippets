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
					StatusCallbackUrl: "https://example.com/your-callback-url",
					Track:             "inbound_track",
				},
			},
		},
	})

	fmt.Print(twiml)
}
