package main

import (
	"fmt"

	"github.com/twilio/twilio-go/twiml"
)

func main() {
	twiml, _ := twiml.Voice([]twiml.Element{
		&twiml.VoicePay{
			Action:           "https://your-callback-function-url.com/pay",
			ChargeAmount:     "10.00",
			PaymentConnector: "My_Pay_Connector",
		},
	})

	fmt.Print(twiml)
}
