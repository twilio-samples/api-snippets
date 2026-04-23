package com.example.usage;

import com.twilio.comms.TwilioCommsClient;

public class Example {
    public static void main(String[] args) {
        TwilioCommsClient client = TwilioCommsClient
            .builder()
            .credentials("<username>", "<password>")
            .build();

        client.pushNotifications().credentials().delete("comms_credential_01h9krwprkeee8fzqspvwy6nq8");
    }
}