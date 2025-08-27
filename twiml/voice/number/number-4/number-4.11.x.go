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
					PhoneNumber:          "+14155555555",
					StatusCallback:       "https://myapp.com/calls/events",
					StatusCallbackEvent:  "initiated ringing answered completed",
					StatusCallbackMethod: "POST",
				},
				&twiml.VoiceNumber{
					PhoneNumber:          "+14153333333",
					StatusCallback:       "https://example.com/calls/events",
					StatusCallbackEvent:  "initiated ringing answered completed",
					StatusCallbackMethod: "POST",
				},
			},
		},
	})

	fmt.Print(twiml)
}
