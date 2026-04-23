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
    request := &twiliocomms.CredentialsUpdateRequest{
        IsDefault: true,
    }
    client.PushNotifications.Credentials.Update(
        context.TODO(),
        "comms_credential_01h9krwprkeee8fzqspvwy6nq8",
        request,
    )
}
