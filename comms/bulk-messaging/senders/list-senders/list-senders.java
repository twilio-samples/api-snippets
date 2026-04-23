package com.example.usage;

import com.twilio.comms.TwilioCommsClient;
import com.twilio.comms.resources.senders.requests.SendersListRequest;
import com.twilio.comms.types.SenderCommunicationChannel;
import com.twilio.comms.types.SenderStatus;
import java.time.OffsetDateTime;

public class Example {
    public static void main(String[] args) {
        TwilioCommsClient client = TwilioCommsClient
            .builder()
            .credentials("<username>", "<password>")
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