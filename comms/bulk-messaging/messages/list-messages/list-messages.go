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
    request := &twiliocomms.MessagesListRequest{
        OperationId: twiliocomms.String(
            "comms_operation_01h9krwprkeee8fzqspvwy6nq8",
        ),
        SessionId: twiliocomms.String(
            "comms_session_01h9krwprkeee8fzqspvwy6nq8",
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
        Profile: twiliocomms.String(
            "mem_profile_01h9krwprkeee8fzqspvwy6nq8",
        ),
        Channel: twiliocomms.MessageSenderChannelSms.Ptr(),
        Status: twiliocomms.MessageStatusQueued.Ptr(),
        Tags: twiliocomms.String(
            "key_1:value;key_2:value;",
        ),
        PageToken: twiliocomms.String(
            "pageToken",
        ),
        PageSize: twiliocomms.Int(
            50,
        ),
    }
    client.Messages.List(
        context.TODO(),
        request,
    )
}
