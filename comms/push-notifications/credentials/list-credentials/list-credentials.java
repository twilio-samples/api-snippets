package com.example.usage;

import com.twilio.comms.TwilioCommsClient;
import com.twilio.comms.resources.pushnotifications.credentials.requests.CredentialsListRequest;
import com.twilio.comms.types.PushCredentialType;
import java.time.OffsetDateTime;

public class Example {
    public static void main(String[] args) {
        TwilioCommsClient client = TwilioCommsClient
            .builder()
            .credentials("<username>", "<password>")
            .build();

        client.pushNotifications().credentials().list(
            CredentialsListRequest
                .builder()
                .startDate(OffsetDateTime.parse("2024-01-15T09:30:00Z"))
                .endDate(OffsetDateTime.parse("2024-01-15T09:30:00Z"))
                .appName("appName")
                .credentialType(PushCredentialType.APN)
                .isDefault(true)
                .pageToken("pageToken")
                .pageSize(50)
                .build()
        );
    }
}