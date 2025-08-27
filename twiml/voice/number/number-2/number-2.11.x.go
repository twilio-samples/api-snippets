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
					PhoneNumber: "858-987-6543",
				},
				&twiml.VoiceNumber{
					PhoneNumber: "415-123-4567",
				},
				&twiml.VoiceNumber{
					PhoneNumber: "619-765-4321",
				},
			},
		},
	})

	fmt.Print(twiml)
}
