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
					PhoneNumber: "415-123-4567",
					Url:         "http://example.com/agent_screen_call",
				},
			},
		},
	})

	fmt.Print(twiml)
}
