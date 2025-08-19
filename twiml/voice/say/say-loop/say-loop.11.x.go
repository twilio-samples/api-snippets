package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceSay{
			Loop:    "2",
			Message: "Hello!",
		},
	})

	fmt.Print(twiml)
}
