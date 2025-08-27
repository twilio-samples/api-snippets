package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceDial{
			InnerElements: []twiml.Element{
				&twiml.VoiceDial{
					Number: "415-123-4567",
				},
				twiml.VoiceSay{
					Message: "Goodbye",
				},
			},
		},
	})

	fmt.Print(twiml)
}
