package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceDial{
			Action: "/handleDialCallStatus",
			Method: "GET",
			Number: "415-123-4567",
		},
		&twiml.VoiceSay{
			Message: "I am unreachable",
		},
	})

	fmt.Print(twiml)
}
