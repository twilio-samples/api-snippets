package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceRefer{
			InnerElements: []twiml.Element{
				&twiml.VoiceSip{
					SipUrl: "sip:alice@example.com",
				},
			},
		},
	})

	fmt.Print(twiml)
}
