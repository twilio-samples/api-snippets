package com.example.usage;

import com.twilio.comms.TwilioCommsClient;
import com.twilio.comms.resources.messages.requests.MessagesListRequest;
import com.twilio.comms.types.MessageSenderChannel;
import com.twilio.comms.types.MessageStatus;
import java.time.OffsetDateTime;

public class Example {
    public static void main(String[] args) {
        TwilioCommsClient client = TwilioCommsClient
            .builder()
            .credentials("<username>", "<password>")
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