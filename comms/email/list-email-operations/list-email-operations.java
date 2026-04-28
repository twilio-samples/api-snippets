package com.example.usage;

import com.twilio.sdk.TwilioComms;
import com.twilio.sdk.resources.emails.requests.EmailsListOperationsRequest;
import com.twilio.sdk.types.CommunicationOperationStatus;
import java.time.OffsetDateTime;

public class Example {
    public static void main(String[] args) {
        TwilioComms client = TwilioComms
            .builder()
            .credentials("TWILIO_ACCOUNT_SID", "TWILIO_AUTH_TOKEN")
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