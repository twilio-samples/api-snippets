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
			PaymentConnector: "My_Generic_Pay_Connector",
			InnerElements: []twiml.Element{
				&twiml.VoiceParameter{
					Name:  "custom_parameter_1",
					Value: "custom_value_1",
				},
			},
		},
	})

	fmt.Print(twiml)
}
