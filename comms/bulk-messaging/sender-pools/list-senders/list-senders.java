package com.example.usage;

import com.twilio.sdk.TwilioComms;
import com.twilio.sdk.resources.senderpools.requests.SenderPoolsListSendersRequest;
import com.twilio.sdk.types.SenderCommunicationChannel;
import com.twilio.sdk.types.SenderStatus;
import java.time.OffsetDateTime;

public class Example {
    public static void main(String[] args) {
        TwilioComms client = TwilioComms
            .builder()
            .credentials("TWILIO_ACCOUNT_SID", "TWILIO_AUTH_TOKEN")
            .build();

        client.senderPools().listSenders(
            "comms_senderpool_01h9krwprkeee8fzqspvwy6nq8",
            SenderPoolsListSendersRequest
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