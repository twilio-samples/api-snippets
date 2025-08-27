package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceGather{
			Action: "/completed",
			Input:  "speech",
			InnerElements: []twiml.Element{
				&twiml.VoiceSay{
					Message: "Welcome to Twilio, please tell us why you're calling",
				},
			},
		},
	})

	fmt.Print(twiml)
}
