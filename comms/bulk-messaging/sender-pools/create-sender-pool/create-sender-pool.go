package example

import (
    context "context"

    twiliocomms "github.com/twilio/twilio-comms-go/twilio"
    client "github.com/twilio/twilio-comms-go/twilio/client"
    option "github.com/twilio/twilio-comms-go/twilio/option"
)

func do() {
    client := client.NewWithOptions(
        option.WithBasicAuth(
            "TWILIO_ACCOUNT_SID",
            "TWILIO_AUTH_TOKEN",
        ),
    )
    request := &twiliocomms.SenderPoolsCreateRequest{
        Name: twiliocomms.String(
            "Sales Leads - APAC",
        ),
        Tags: &twiliocomms.Tags{
            "region": "APAC",
        },
    }
    client.SenderPools.Create(
        context.TODO(),
        request,
    )
}
