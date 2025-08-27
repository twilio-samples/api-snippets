package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceRedirect{
			Method: "POST",
			Url:    "http://pigeons.com/twiml.xml",
		},
	})

	fmt.Print(twiml)
}
