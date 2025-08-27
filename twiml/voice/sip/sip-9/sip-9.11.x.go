package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceDial{
			Action:       "/handle_post_dial",
			CallerId:     "bob",
			HangupOnStar: "true",
			Method:       "POST",
			Record:       "record-from-answer",
			Timeout:      "10",
			InnerElements: []twiml.Element{
				&twiml.VoiceSip{
					Method: "POST",
					SipUrl: "sip:kate@example.com?x-customheader=foo",
					Url:    "/handle_screening_on_answer",
				},
			},
		},
	})

	fmt.Print(twiml)
}
