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
    client.Messages.FetchOperation(
        context.TODO(),
        "comms_operation_01h9krwprkeee8fzqspvwy6nq8",
    )
}
