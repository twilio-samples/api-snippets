package com.example.usage;

import com.twilio.sdk.TwilioComms;
import com.twilio.sdk.resources.senders.requests.SendersSearchRequest;
import com.twilio.sdk.types.SenderCommunicationChannel;

public class Example {
    public static void main(String[] args) {
        TwilioComms client = TwilioComms
            .builder()
            .credentials("TWILIO_ACCOUNT_SID", "TWILIO_AUTH_TOKEN")
            .build();

        client.senders().search(
            SendersSearchRequest
                .builder()
                .address("+14153902337")
                .channel(SenderCommunicationChannel.SMS)
                .build()
        );
    }
}