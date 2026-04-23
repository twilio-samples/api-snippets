package com.example.usage;

import com.twilio.comms.TwilioCommsClient;
import com.twilio.comms.resources.messages.requests.MessagesSendRequest;
import com.twilio.comms.resources.messages.types.MessagesSendRequestContent;
import com.twilio.comms.resources.messages.types.MessagesSendRequestToItem;
import com.twilio.comms.resources.messages.types.MessagesSendRequestToItemAddress;
import com.twilio.comms.types.MessageAddressChannel;
import com.twilio.comms.types.MessageContentTextWithMedia;
import java.util.Arrays;
import java.util.Optional;

public class Example {
    public static void main(String[] args) {
        TwilioCommsClient client = TwilioCommsClient
            .builder()
            .credentials("<username>", "<password>")
            .build();

        client.messages().send(
            MessagesSendRequest
                .builder()
                .content(
                    MessagesSendRequestContent.of(
                        MessageContentTextWithMedia
                            .builder()
                            .text(Optional.of("Hello, World!"))
                            .build()
                    )
                )
                .to(
                    Arrays.asList(
                        MessagesSendRequestToItem.of(
                            MessagesSendRequestToItemAddress
                                .builder()
                                .address("+14153902337")
                                .channel(MessageAddressChannel.PHONE)
                                .build()
                        )
                    )
                )
                .build()
        );
    }
}