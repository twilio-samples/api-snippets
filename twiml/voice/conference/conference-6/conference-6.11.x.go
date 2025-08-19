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
					Beep:                   "false",
					EndConferenceOnExit:    "true",
					Name:                   "NoMusicNoBeepRoom",
					StartConferenceOnEnter: "true",
					WaitUrl:                "http://your-webhook-host.com",
				},
			},
		},
	})

	fmt.Print(twiml)
}
