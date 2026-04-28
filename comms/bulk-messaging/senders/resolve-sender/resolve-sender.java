package com.example.usage;

import com.twilio.sdk.TwilioComms;
import com.twilio.sdk.resources.senders.types.SendersResolveRequest;
import com.twilio.sdk.resources.senders.types.SendersResolveRequestSenderPoolId;
import com.twilio.sdk.types.CommunicationChannel;
import com.twilio.sdk.types.CommunicationRecipient;
import java.util.Arrays;

public class Example {
    public static void main(String[] args) {
        TwilioComms client = TwilioComms
            .builder()
            .credentials("TWILIO_ACCOUNT_SID", "TWILIO_AUTH_TOKEN")
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