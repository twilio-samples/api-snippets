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
    request := &twiliocomms.SenderPoolsListOperationsRequest{
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
        Status: twiliocomms.OperationStatusProcessing.Ptr(),
        PageToken: twiliocomms.String(
            "pageToken",
        ),
        PageSize: twiliocomms.Int(
            50,
        ),
    }
    client.SenderPools.ListOperations(
        context.TODO(),
        request,
    )
}
