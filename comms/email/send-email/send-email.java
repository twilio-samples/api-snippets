package com.example.usage;

import com.twilio.comms.TwilioCommsClient;
import com.twilio.comms.resources.emails.requests.EmailsSendRequest;
import com.twilio.comms.resources.emails.types.EmailsSendRequestFrom;
import com.twilio.comms.resources.emails.types.EmailsSendRequestToItem;
import com.twilio.comms.resources.emails.types.EmailsSendRequestToItemAddress;
import com.twilio.comms.types.EmailAddressSender;
import com.twilio.comms.types.EmailAttachmentsItem;
import com.twilio.comms.types.EmailHtmlContent;
import java.util.Arrays;
import java.util.Optional;

public class Example {
    public static void main(String[] args) {
        TwilioCommsClient client = TwilioCommsClient
            .builder()
            .credentials("<username>", "<password>")
            .build();

        client.emails().send(
            EmailsSendRequest
                .builder()
                .from(
                    EmailsSendRequestFrom.of(
                        EmailAddressSender
                            .builder()
                            .address("support@example.company.io")
                            .name("Cool Co Support")
                            .build()
                    )
                )
                .content(
                    EmailHtmlContent
                        .builder()
                        .html("<html><body>Hey, <br/><br/>Cake</b></body></html>")
                        .subject("Re: Wedding Cake")
                        .text("Hey, the cake is ready.")
                        .attachments(
                            Arrays.asList(
                                EmailAttachmentsItem
                                    .builder()
                                    .filename("filename")
                                    .contentType("contentType")
                                    .content("content")
                                    .build()
                            )
                        )
                        .build()
                )
                .to(
                    Arrays.asList(
                        EmailsSendRequestToItem.of(
                            EmailsSendRequestToItemAddress
                                .builder()
                                .address("bob@example.com")
                                .name(Optional.of("Bob Smith"))
                                .build()
                        )
                    )
                )
                .build()
        );
    }
}