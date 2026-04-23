package com.example.usage;

import com.twilio.comms.TwilioCommsClient;
import com.twilio.comms.resources.senderpools.types.SenderPoolsAddSenderRequestItem;
import java.util.Arrays;

public class Example {
    public static void main(String[] args) {
        TwilioCommsClient client = TwilioCommsClient
            .builder()
            .credentials("<username>", "<password>")
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