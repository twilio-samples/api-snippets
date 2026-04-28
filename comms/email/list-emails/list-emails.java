package com.example.usage;

import com.twilio.sdk.TwilioComms;
import com.twilio.sdk.resources.emails.requests.EmailsListRequest;
import com.twilio.sdk.types.EmailStatus;
import java.time.OffsetDateTime;

public class Example {
    public static void main(String[] args) {
        TwilioComms client = TwilioComms
            .builder()
            .credentials("TWILIO_ACCOUNT_SID", "TWILIO_AUTH_TOKEN")
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