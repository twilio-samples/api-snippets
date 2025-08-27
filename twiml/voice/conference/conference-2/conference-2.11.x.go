package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceDial{
			InnerElements: []twiml.Element{
				&twiml.VoiceConference{
					Name:                   "moderated-conference-room",
					StartConferenceOnEnter: "false",
				},
			},
		},
	})

	fmt.Print(twiml)
}
