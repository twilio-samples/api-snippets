package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceDial{
			InnerElements: []twiml.Element{
				&twiml.VoiceSip{
					SipUrl: "kate@example.com",
				},
			},
		},
	})

	fmt.Print(twiml)
}
