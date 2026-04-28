package com.example.usage;

import com.twilio.sdk.TwilioComms;
import com.twilio.sdk.resources.emails.requests.EmailsSendRequest;
import com.twilio.sdk.resources.emails.types.EmailsSendRequestFrom;
import com.twilio.sdk.resources.emails.types.EmailsSendRequestToItem;
import com.twilio.sdk.resources.emails.types.EmailsSendRequestToItemAddress;
import com.twilio.sdk.types.EmailAddressSender;
import com.twilio.sdk.types.EmailAttachmentsItem;
import com.twilio.sdk.types.EmailHtmlContent;
import java.util.Arrays;
import java.util.Optional;

public class Example {
    public static void main(String[] args) {
        TwilioComms client = TwilioComms
            .builder()
            .credentials("TWILIO_ACCOUNT_SID", "TWILIO_AUTH_TOKEN")
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
                        .html("<html><body>Hey, <br/><br/><b>Cake</b></body></html>")
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