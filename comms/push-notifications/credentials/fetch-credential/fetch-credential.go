package example

import (
    context "context"

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
    client.PushNotifications.Credentials.Fetch(
        context.TODO(),
        "comms_credential_01h9krwprkeee8fzqspvwy6nq8",
    )
}
