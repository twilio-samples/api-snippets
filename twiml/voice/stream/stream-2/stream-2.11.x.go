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
					Url: "wss://mystream.ngrok.io/example",
					InnerElements: []twiml.Element{
						&twiml.VoiceParameter{
							Name:  "FirstName",
							Value: "Jane",
						},
						&twiml.VoiceParameter{
							Name:  "LastName",
							Value: "Doe",
						},
					},
				},
			},
		},
	})

	fmt.Print(twiml)
}
