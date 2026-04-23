package com.example.usage;

import com.twilio.comms.TwilioCommsClient;
import com.twilio.comms.resources.pushnotifications.users.requests.UsersListRequest;

public class Example {
    public static void main(String[] args) {
        TwilioCommsClient client = TwilioCommsClient
            .builder()
            .credentials("<username>", "<password>")
            .build();

        client.pushNotifications().users().list(
            UsersListRequest
                .builder()
                .pageToken("pageToken")
                .pageSize(50)
                .build()
        );
    }
}