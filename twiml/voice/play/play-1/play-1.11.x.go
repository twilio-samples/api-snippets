package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoicePlay{
			Loop: "10",
			Url:  "https://api.twilio.com/cowbell.mp3",
		},
	})

	fmt.Print(twiml)
}
