package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceDial{
			InnerElements: []twiml.Element{
				&twiml.VoiceQueue{
					Url:  "about_to_connect.xml",
					Name: "support",
				},
			},
		},
	})

	fmt.Print(twiml)
}
