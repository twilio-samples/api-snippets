package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceDial{
			Sequential: "true",
			InnerElements: []twiml.Element{
				&twiml.VoiceSip{
					SipUrl: "sip:alice@example.com",
				},
				&twiml.VoiceSip{
					SipUrl: "sip:bob@example.com",
				},
				&twiml.VoiceSip{
					SipUrl: "sip:charlie@example.com",
				},
			},
		},
	})

	fmt.Print(twiml)
}
