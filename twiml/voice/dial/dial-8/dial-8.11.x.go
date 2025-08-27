package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceDial{
			ReferUrl: "https://example.com/handler",
			InnerElements: []twiml.Element{
				&twiml.VoiceSip{
					SipUrl: "sip:AgentA@xyz.sip.us1.twilio.com?User-to-User=123456789%3Bencoding%3Dhex&amp;X-Name=Agent%2C+A",
				},
			},
		},
	})

	fmt.Print(twiml)
}
