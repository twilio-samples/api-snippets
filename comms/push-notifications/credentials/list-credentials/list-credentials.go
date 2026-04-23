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
            "<username>",
            "<password>",
        ),
    )
    request := &twiliocomms.CredentialsListRequest{
        StartDate: twiliocomms.Time(
            twiliocomms.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        EndDate: twiliocomms.Time(
            twiliocomms.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        AppName: twiliocomms.String(
            "appName",
        ),
        CredentialType: twiliocomms.PushCredentialTypeApn.Ptr(),
        IsDefault: twiliocomms.Bool(
            true,
        ),
        PageToken: twiliocomms.String(
            "pageToken",
        ),
        PageSize: twiliocomms.Int(
            50,
        ),
    }
    client.PushNotifications.Credentials.List(
        context.TODO(),
        request,
    )
}
