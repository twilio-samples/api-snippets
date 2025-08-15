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
					InnerElements: []twiml.Element{
						&twiml.VoiceParameter{
							Name:  "foo",
							Value: "bar",
						},
						&twiml.VoiceParameter{
							Name:  "hint",
							Value: "Annoyed customer",
						},
					},
				},
			},
		},
	})

	fmt.Print(twiml)
}
