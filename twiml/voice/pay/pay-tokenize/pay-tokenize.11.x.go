package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoicePay{
			ChargeAmount: "0",
			TokenType:    "one-time",
		},
	})

	fmt.Print(twiml)
}
