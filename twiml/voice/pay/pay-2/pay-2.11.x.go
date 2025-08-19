package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoiceSay{
			Message: "Calling Twilio Pay",
		},
		&twiml.VoicePay{
			ChargeAmount: "20.45",
		},
	})

	fmt.Print(twiml)
}
