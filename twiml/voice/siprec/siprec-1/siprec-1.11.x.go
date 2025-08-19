package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceStart{
			InnerElements: []twiml.Element{
				&twiml.VoiceSipRec{
					ConnectorName: "Gridspace_1",
					Name:          "My SIPREC Stream",
				},
			},
		},
	})

	fmt.Print(twiml)
}
