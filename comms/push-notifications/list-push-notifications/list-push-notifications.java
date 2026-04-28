package com.example.usage;

import com.twilio.sdk.TwilioComms;
import com.twilio.sdk.resources.pushnotifications.requests.PushNotificationsListRequest;
import com.twilio.sdk.types.PushNotificationProvider;
import com.twilio.sdk.types.PushNotificationStatus;
import java.time.OffsetDateTime;

public class Example {
    public static void main(String[] args) {
        TwilioComms client = TwilioComms
            .builder()
            .credentials("TWILIO_ACCOUNT_SID", "TWILIO_AUTH_TOKEN")
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