package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoicePause{
			Length: "5",
		},
		&twiml.VoiceSay{
			Message: "Hi there.",
		},
	})

	fmt.Print(twiml)
}
