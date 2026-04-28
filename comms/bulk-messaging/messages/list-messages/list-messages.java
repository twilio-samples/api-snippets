package com.example.usage;

import com.twilio.sdk.TwilioComms;
import com.twilio.sdk.resources.messages.requests.MessagesListRequest;
import com.twilio.sdk.types.MessageSenderChannel;
import com.twilio.sdk.types.MessageStatus;
import java.time.OffsetDateTime;

public class Example {
    public static void main(String[] args) {
        TwilioComms client = TwilioComms
            .builder()
            .credentials("TWILIO_ACCOUNT_SID", "TWILIO_AUTH_TOKEN")
            .build();

        client.messages().list(
            MessagesListRequest
                .builder()
                .operationId("comms_operation_01h9krwprkeee8fzqspvwy6nq8")
                .sessionId("comms_session_01h9krwprkeee8fzqspvwy6nq8")
                .startDate(OffsetDateTime.parse("2024-01-15T09:30:00Z"))
                .endDate(OffsetDateTime.parse("2024-01-15T09:30:00Z"))
                .profile("mem_profile_01h9krwprkeee8fzqspvwy6nq8")
                .channel(MessageSenderChannel.SMS)
                .status(MessageStatus.QUEUED)
                .tags("key_1:value;key_2:value;")
                .pageToken("pageToken")
                .pageSize(50)
                .build()
        );
    }
}