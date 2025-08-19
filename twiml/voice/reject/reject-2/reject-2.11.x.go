package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceReject{
			Reason: "busy",
		},
	})

	fmt.Print(twiml)
}
