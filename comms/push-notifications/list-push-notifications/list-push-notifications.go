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
    request := &twiliocomms.PushNotificationsListRequest{
        OperationId: twiliocomms.String(
            "comms_operation_01h9krwprkeee8fzqspvwy6nq8",
        ),
        Status: twiliocomms.PushNotificationStatusScheduled.Ptr(),
        Tags: twiliocomms.String(
            "key_1:value;key_2:value;",
        ),
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
        User: twiliocomms.String(
            "comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8",
        ),
        Provider: twiliocomms.PushNotificationProviderApn.Ptr(),
        PageToken: twiliocomms.String(
            "pageToken",
        ),
        PageSize: twiliocomms.Int(
            50,
        ),
    }
    client.PushNotifications.List(
        context.TODO(),
        request,
    )
}
