package com.example.usage;

import com.twilio.comms.TwilioCommsClient;
import com.twilio.comms.resources.senders.types.SendersResolveRequest;
import com.twilio.comms.resources.senders.types.SendersResolveRequestSenderPoolId;
import com.twilio.comms.types.CommunicationChannel;
import com.twilio.comms.types.CommunicationRecipient;
import java.util.Arrays;

public class Example {
    public static void main(String[] args) {
        TwilioCommsClient client = TwilioCommsClient
            .builder()
            .credentials("<username>", "<password>")
            .build();

        client.senders().resolve(
            SendersResolveRequest.of(
                SendersResolveRequestSenderPoolId
                    .builder()
                    .recipientAddresses(
                        Arrays.asList(
                            CommunicationRecipient
                                .builder()
                                .address("+14153902337")
                                .channel(CommunicationChannel.PHONE)
                                .build(),
                            CommunicationRecipient
                                .builder()
                                .address("+14153902337")
                                .channel(CommunicationChannel.WHATSAPP)
                                .build(),
                            CommunicationRecipient
                                .builder()
                                .address("davidpletnjov@example.com")
                                .channel(CommunicationChannel.EMAIL)
                                .build()
                        )
                    )
                    .build()
            )
        );
    }
}