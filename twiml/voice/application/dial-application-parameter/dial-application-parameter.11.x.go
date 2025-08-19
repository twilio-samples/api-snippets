package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceDial{
			InnerElements: []twiml.Element{
				&twiml.VoiceApplication{
					InnerElements: []twiml.Element{
						&twiml.VoiceApplicationSid{
							Sid: "AP1234567890abcdef1234567890abcd",
						},
						&twiml.VoiceParameter{
							Name:  "AccountNumber",
							Value: "12345",
						},
						&twiml.VoiceParameter{
							Name:  "TicketNumber",
							Value: "9876",
						},
					},
				},
			},
		},
	})

	fmt.Print(twiml)
}
