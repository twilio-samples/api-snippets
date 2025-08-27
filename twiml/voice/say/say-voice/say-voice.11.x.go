package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceSay{
			Language: "fr-FR",
			Message:  "Bonjour Je m'appelle Mathieu.",
			Voice:    "Polly.Mathieu",
		},
	})

	fmt.Print(twiml)
}
