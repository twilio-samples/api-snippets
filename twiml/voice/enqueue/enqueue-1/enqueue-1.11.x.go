package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceEnqueue{
			Name:    "support",
			WaitUrl: "wait-music.xml",
		},
	})

	fmt.Print(twiml)
}
