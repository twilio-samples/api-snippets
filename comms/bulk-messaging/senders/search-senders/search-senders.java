package com.example.usage;

import com.twilio.comms.TwilioCommsClient;
import com.twilio.comms.resources.senders.requests.SendersSearchRequest;
import com.twilio.comms.types.SenderCommunicationChannel;

public class Example {
    public static void main(String[] args) {
        TwilioCommsClient client = TwilioCommsClient
            .builder()
            .credentials("<username>", "<password>")
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