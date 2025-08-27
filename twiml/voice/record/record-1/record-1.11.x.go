package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceRecord{
			Timeout:    "10",
			Transcribe: "true",
		},
	})

	fmt.Print(twiml)
}
