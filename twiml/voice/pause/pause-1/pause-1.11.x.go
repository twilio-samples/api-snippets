package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceSay{
			Message: "I will pause 10 seconds starting now!",
		},
		&twiml.VoicePause{
			Length: "10",
		},
		&twiml.VoiceSay{
			Message: "I just paused 10 seconds",
		},
	})

	fmt.Print(twiml)
}
