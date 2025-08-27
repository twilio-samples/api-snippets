package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceDial{
			InnerElements: []twiml.Element{
				&twiml.VoiceSip{
					Password: "1234",
					SipUrl:   "kate@example.com",
					Username: "admin",
				},
			},
		},
	})

	fmt.Print(twiml)
}
