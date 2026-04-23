package com.example.usage;

import com.twilio.comms.TwilioCommsClient;
import com.twilio.comms.resources.pushnotifications.deviceregistrations.requests.DeviceRegistrationsListRequest;

public class Example {
    public static void main(String[] args) {
        TwilioCommsClient client = TwilioCommsClient
            .builder()
            .credentials("<username>", "<password>")
            .build();

        client.pushNotifications().deviceRegistrations().list(
            DeviceRegistrationsListRequest
                .builder()
                .userId("comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8")
                .appName("appName")
                .pageToken("pageToken")
                .pageSize(50)
                .build()
        );
    }
}