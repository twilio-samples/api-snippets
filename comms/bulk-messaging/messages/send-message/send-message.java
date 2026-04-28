package com.example.usage;

import com.twilio.sdk.TwilioComms;
import com.twilio.sdk.resources.messages.requests.MessagesSendRequest;
import com.twilio.sdk.resources.messages.types.MessagesSendRequestContent;
import com.twilio.sdk.resources.messages.types.MessagesSendRequestFrom;
import com.twilio.sdk.resources.messages.types.MessagesSendRequestToItem;
import com.twilio.sdk.resources.messages.types.MessagesSendRequestToItemAddress;
import com.twilio.sdk.types.MessageAddressChannel;
import com.twilio.sdk.types.MessageAddressSender;
import com.twilio.sdk.types.MessageContentTextWithMedia;
import com.twilio.sdk.types.MessageSenderChannel;
import java.util.Arrays;
import java.util.Optional;

public class Example {
    public static void main(String[] args) {
        TwilioComms client = TwilioComms
            .builder()
            .credentials("TWILIO_ACCOUNT_SID", "TWILIO_AUTH_TOKEN")
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
                                .address("+12065558844")
                                .channel(MessageAddressChannel.PHONE)
                                .build()
                        )
                    )
                )
                .from(
                    MessagesSendRequestFrom.of(
                        MessageAddressSender
                            .builder()
                            .address("+14153901002")
                            .channel(MessageSenderChannel.SMS)
                            .build()
                    )
                )
                .build()
        );
    }
}