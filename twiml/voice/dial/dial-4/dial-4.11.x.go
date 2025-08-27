package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceDial{
			CallerId: "+15551112222",
			InnerElements: []twiml.Element{
				&twiml.VoiceNumber{
					PhoneNumber: "+15558675310",
				},
			},
		},
	})

	fmt.Print(twiml)
}
