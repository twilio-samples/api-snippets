package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceDial{
			InnerElements: []twiml.Element{
				&twiml.VoiceConference{
					Name:                "EventedConf",
					StatusCallback:      "https://myapp.com/events",
					StatusCallbackEvent: "start end join leave mute hold",
				},
			},
		},
	})

	fmt.Print(twiml)
}
