package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceEnqueue{
			Name:    "support",
			WaitUrl: "wait.xml",
		},
		&twiml.VoiceSay{
			Message: "Unfortunately, the support line has closed. Please call again tomorrow.",
		},
	})

	fmt.Print(twiml)
}
