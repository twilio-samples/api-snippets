package com.example.usage;

import com.twilio.comms.TwilioCommsClient;
import com.twilio.comms.resources.pushnotifications.credentials.requests.CredentialsUpdateRequest;

public class Example {
    public static void main(String[] args) {
        TwilioCommsClient client = TwilioCommsClient
            .builder()
            .credentials("<username>", "<password>")
            .build();

        client.pushNotifications().credentials().update(
            "comms_credential_01h9krwprkeee8fzqspvwy6nq8",
            CredentialsUpdateRequest
                .builder()
                .isDefault(true)
                .build()
        );
    }
}