package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceDial{
			CallerId: "+1888XXXXXXX",
			InnerElements: []twiml.Element{
				&twiml.VoiceNumber{
					PhoneNumber: "858-987-6543",
				},
				&twiml.VoiceClient{
					Identity: "joey",
				},
				&twiml.VoiceClient{
					Identity: "charlie",
				},
			},
		},
	})

	fmt.Print(twiml)
}
