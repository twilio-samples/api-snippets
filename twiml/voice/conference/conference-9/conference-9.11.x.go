package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceDial{
			Action:       "handleLeaveConference.php",
			HangupOnStar: "true",
			Method:       "POST",
			TimeLimit:    "30",
			InnerElements: []twiml.Element{
				&twiml.VoiceConference{
					Name: "LoveTwilio",
				},
			},
		},
	})

	fmt.Print(twiml)
}
