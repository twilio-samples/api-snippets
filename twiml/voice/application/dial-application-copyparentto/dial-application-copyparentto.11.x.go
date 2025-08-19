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
					CopyParentTo: "true",
					InnerElements: []twiml.Element{
						&twiml.VoiceApplicationSid{
							Sid: "AP1234567890abcdef1234567890abcd",
						},
					},
				},
			},
		},
	})

	fmt.Print(twiml)
}
