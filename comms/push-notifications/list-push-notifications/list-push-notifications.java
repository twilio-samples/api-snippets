package com.example.usage;

import com.twilio.comms.TwilioCommsClient;
import com.twilio.comms.resources.pushnotifications.requests.PushNotificationsListRequest;
import com.twilio.comms.types.PushNotificationProvider;
import com.twilio.comms.types.PushNotificationStatus;
import java.time.OffsetDateTime;

public class Example {
    public static void main(String[] args) {
        TwilioCommsClient client = TwilioCommsClient
            .builder()
            .credentials("<username>", "<password>")
            .build();

        client.pushNotifications().list(
            PushNotificationsListRequest
                .builder()
                .operationId("comms_operation_01h9krwprkeee8fzqspvwy6nq8")
                .status(PushNotificationStatus.SCHEDULED)
                .tags("key_1:value;key_2:value;")
                .startDate(OffsetDateTime.parse("2024-01-15T09:30:00Z"))
                .endDate(OffsetDateTime.parse("2024-01-15T09:30:00Z"))
                .user("comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8")
                .provider(PushNotificationProvider.APN)
                .pageToken("pageToken")
                .pageSize(50)
                .build()
        );
    }
}