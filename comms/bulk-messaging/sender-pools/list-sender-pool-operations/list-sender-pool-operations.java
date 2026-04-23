package com.example.usage;

import com.twilio.comms.TwilioCommsClient;
import com.twilio.comms.resources.senderpools.requests.SenderPoolsListOperationsRequest;
import com.twilio.comms.types.OperationStatus;
import java.time.OffsetDateTime;

public class Example {
    public static void main(String[] args) {
        TwilioCommsClient client = TwilioCommsClient
            .builder()
            .credentials("<username>", "<password>")
            .build();

        client.senderPools().listOperations(
            SenderPoolsListOperationsRequest
                .builder()
                .startDate(OffsetDateTime.parse("2024-01-15T09:30:00Z"))
                .endDate(OffsetDateTime.parse("2024-01-15T09:30:00Z"))
                .status(OperationStatus.PROCESSING)
                .pageToken("pageToken")
                .pageSize(50)
                .build()
        );
    }
}