package com.example.usage;

import com.twilio.comms.TwilioCommsClient;
import com.twilio.comms.resources.messages.requests.MessagesListOperationsRequest;
import com.twilio.comms.types.CommunicationOperationStatus;
import java.time.OffsetDateTime;

public class Example {
    public static void main(String[] args) {
        TwilioCommsClient client = TwilioCommsClient
            .builder()
            .credentials("<username>", "<password>")
            .build();

        client.messages().listOperations(
            MessagesListOperationsRequest
                .builder()
                .startDate(OffsetDateTime.parse("2024-01-15T09:30:00Z"))
                .endDate(OffsetDateTime.parse("2024-01-15T09:30:00Z"))
                .status(CommunicationOperationStatus.SCHEDULED)
                .pageToken("pageToken")
                .pageSize(50)
                .build()
        );
    }
}