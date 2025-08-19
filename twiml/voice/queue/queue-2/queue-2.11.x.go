package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceDial{
			InnerElements: []twiml.Element{
				&twiml.VoiceSay{
					Message: "You will now be connected to an agent.",
				},
			},
		},
	})

	fmt.Print(twiml)
}
