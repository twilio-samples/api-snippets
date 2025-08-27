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
					SipUrl: "sip:kate@example.com?x-mycustomheader=foo&amp;x-myotherheader=bar",
				},
			},
		},
	})

	fmt.Print(twiml)
}
