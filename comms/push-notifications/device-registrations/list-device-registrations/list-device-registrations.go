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
    request := &twiliocomms.DeviceRegistrationsListRequest{
        UserId: "comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8",
        AppName: twiliocomms.String(
            "appName",
        ),
        PageToken: twiliocomms.String(
            "pageToken",
        ),
        PageSize: twiliocomms.Int(
            50,
        ),
    }
    client.PushNotifications.DeviceRegistrations.List(
        context.TODO(),
        request,
    )
}
