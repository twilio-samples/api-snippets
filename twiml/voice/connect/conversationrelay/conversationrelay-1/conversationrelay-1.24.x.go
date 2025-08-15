package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceConnect{
			InnerElements: []twiml.Element{
				&twiml.VoiceConversationRelay{
					Url: "wss://mywebsocketserver.com/websocket",
					InnerElements: []twiml.Element{
						&twiml.VoiceLanguage{
							Code:                  "sv-SE",
							TtsProvider:           "amazon",
							Voice:                 "Elin-Neural",
							TranscriptionProvider: "google",
							SpeechModel:           "long",
						},
						&twiml.VoiceLanguage{
							Code:        "en-US",
							TtsProvider: "google",
							Voice:       "en-US-Journey-0",
						},
					},
				},
			},
		},
	})

	fmt.Print(twiml)
}
