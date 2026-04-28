package com.example.usage;

import com.twilio.sdk.TwilioComms;
import com.twilio.sdk.resources.senders.requests.SendersListRequest;
import com.twilio.sdk.types.SenderCommunicationChannel;
import com.twilio.sdk.types.SenderStatus;
import java.time.OffsetDateTime;

public class Example {
    public static void main(String[] args) {
        TwilioComms client = TwilioComms
            .builder()
            .credentials("TWILIO_ACCOUNT_SID", "TWILIO_AUTH_TOKEN")
            .build();

        client.senders().list(
            SendersListRequest
                .builder()
                .channel(SenderCommunicationChannel.SMS)
                .status(SenderStatus.ACTIVATED)
                .startDate(OffsetDateTime.parse("2024-01-15T09:30:00Z"))
                .endDate(OffsetDateTime.parse("2024-01-15T09:30:00Z"))
                .pageToken("pageToken")
                .pageSize(50)
                .build()
        );
    }
}