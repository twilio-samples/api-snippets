package main

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go/client"
)

func main() {
	// Your Auth Token from twilio.com/console
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	// Initialize the request validator
	requestValidator := client.NewRequestValidator(authToken)

	// Store Twilio's request URL (the url of your webhook) as a variable
	// including all query parameters
	url := "https://example.com/myapp?bodySHA256=5ccde7145dfb8f56479710896586cb9d5911809d83afbe34627818790db0aec9"

	// Store the application/json body from Twilio's request as a variable
	// In practice, this MUST include all received parameters, not a
	// hardcoded list of parameters that you receive today. New parameters
	// may be added without notice.
	body :=  []byte("{\"CallSid\":\"CA1234567890ABCDE\",\"Caller\":\"+12349013030\"}")

	// Store the X-Twilio-Signature header attached to the request as a variable
	signature := "hqeF3G9Hrnv6/R0jOhoYDD2PPUs="

	// Check if the incoming signature is valid for your application URL and the incoming parameters
	fmt.Println(requestValidator.ValidateBody(url, body, signature))
}
