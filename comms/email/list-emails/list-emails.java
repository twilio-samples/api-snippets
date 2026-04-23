package com.example.usage;

import com.twilio.comms.TwilioCommsClient;
import com.twilio.comms.resources.emails.requests.EmailsListRequest;
import com.twilio.comms.types.EmailStatus;
import java.time.OffsetDateTime;

public class Example {
    public static void main(String[] args) {
        TwilioCommsClient client = TwilioCommsClient
            .builder()
            .credentials("<username>", "<password>")
            .build();

        client.emails().list(
            EmailsListRequest
                .builder()
                .operationId("comms_operation_01h9krwprkeee8fzqspvwy6nq8")
                .startDate(OffsetDateTime.parse("2024-01-15T09:30:00Z"))
                .endDate(OffsetDateTime.parse("2024-01-15T09:30:00Z"))
                .status(EmailStatus.SCHEDULED)
                .tags("key_1:value;key_2:value;")
                .pageToken("pageToken")
                .pageSize(50)
                .build()
        );
    }
}