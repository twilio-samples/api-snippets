package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceDial{
			InnerElements: []twiml.Element{
				&twiml.VoiceNumber{
					PhoneNumber: "415-123-4567",
					SendDigits:  "www1928",
				},
			},
		},
	})

	fmt.Print(twiml)
}
