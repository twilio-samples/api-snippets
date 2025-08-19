package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceDial{
			InnerElements: []twiml.Element{
				&twiml.VoiceClient{
					InnerElements: []twiml.Element{
						&twiml.VoiceIdentity{
							ClientIdentity: "user_jane",
						},
						&twiml.VoiceParameter{
							Name:  "FirstName",
							Value: "Jane",
						},
						&twiml.VoiceParameter{
							Name:  "LastName",
							Value: "Doe",
						},
					},
				},
			},
		},
	})

	fmt.Print(twiml)
}
