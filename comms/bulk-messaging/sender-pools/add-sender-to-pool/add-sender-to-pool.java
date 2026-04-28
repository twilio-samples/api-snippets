package com.example.usage;

import com.twilio.sdk.TwilioComms;
import com.twilio.sdk.resources.senderpools.types.SenderPoolsAddSenderRequestItem;
import java.util.Arrays;

public class Example {
    public static void main(String[] args) {
        TwilioComms client = TwilioComms
            .builder()
            .credentials("TWILIO_ACCOUNT_SID", "TWILIO_AUTH_TOKEN")
            .build();

        client.senderPools().addSender(
            "comms_senderpool_01h9krwprkeee8fzqspvwy6nq8",
            Arrays.asList(
                SenderPoolsAddSenderRequestItem
                    .builder()
                    .senderId("comms_sender_01h9krwprkeee8fzqspvwy6nq8")
                    .build()
            )
        );
    }
}