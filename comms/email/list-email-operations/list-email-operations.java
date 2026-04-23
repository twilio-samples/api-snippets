package com.example.usage;

import com.twilio.comms.TwilioCommsClient;
import com.twilio.comms.resources.emails.requests.EmailsListOperationsRequest;
import com.twilio.comms.types.CommunicationOperationStatus;
import java.time.OffsetDateTime;

public class Example {
    public static void main(String[] args) {
        TwilioCommsClient client = TwilioCommsClient
            .builder()
            .credentials("<username>", "<password>")
            .build();

        client.emails().listOperations(
            EmailsListOperationsRequest
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