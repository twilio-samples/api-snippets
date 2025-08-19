package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoicePay{
			PaymentConnector: "Stripe_Connector_1",
		},
	})

	fmt.Print(twiml)
}
