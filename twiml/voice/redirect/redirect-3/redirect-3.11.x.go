package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceRedirect{
			Url: "../nextInstructions",
		},
	})

	fmt.Print(twiml)
}
