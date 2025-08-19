package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceConnect{
			InnerElements: []twiml.Element{
				&twiml.VoiceStream{
					Url: "wss://example.com/audiostream",
				},
			},
		},
		&twiml.VoiceSay{
			Message: "This TwiML instruction is unreachable unless the Stream is ended by your WebSocket server.",
		},
	})

	fmt.Print(twiml)
}
