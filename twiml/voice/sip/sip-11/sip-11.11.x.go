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
					PhoneNumber: "+12143211432",
				},
				&twiml.VoiceSip{
					SipUrl: "sip:alice-soft-phone@example.com",
				},
				&twiml.VoiceSip{
					SipUrl: "sip:alice-desk-phone@example.com",
				},
				&twiml.VoiceSip{
					SipUrl: "sip:alice-mobile-client@example.com",
				},
			},
		},
	})

	fmt.Print(twiml)
}
