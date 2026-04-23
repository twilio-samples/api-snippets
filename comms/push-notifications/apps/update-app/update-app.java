package com.example.usage;

import com.twilio.comms.TwilioCommsClient;
import com.twilio.comms.resources.pushnotifications.apps.requests.AppsUpdateRequest;

public class Example {
    public static void main(String[] args) {
        TwilioCommsClient client = TwilioCommsClient
            .builder()
            .credentials("<username>", "<password>")
            .build();

        client.pushNotifications().apps().update(
            "appName",
            AppsUpdateRequest
                .builder()
                .isDefault(true)
                .build()
        );
    }
}