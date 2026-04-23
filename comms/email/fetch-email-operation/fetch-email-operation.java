package com.example.usage;

import com.twilio.comms.TwilioCommsClient;

public class Example {
    public static void main(String[] args) {
        TwilioCommsClient client = TwilioCommsClient
            .builder()
            .credentials("<username>", "<password>")
            .build();

        client.emails().fetchOperation("comms_operation_01h9krwprkeee8fzqspvwy6nq8");
    }
}