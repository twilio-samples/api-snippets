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
					Identity:             "joey",
					StatusCallback:       "https://myapp.com/calls/events",
					StatusCallbackEvent:  "initiated ringing answered completed",
					StatusCallbackMethod: "POST",
				},
			},
		},
	})

	fmt.Print(twiml)
}
