package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceGather{
			Input:     "speech dtmf",
			NumDigits: "1",
			Timeout:   "3",
			InnerElements: []twiml.Element{
				&twiml.VoiceSay{
					Message: "Please press 1 or say sales for sales.",
				},
			},
		},
	})

	fmt.Print(twiml)
}
