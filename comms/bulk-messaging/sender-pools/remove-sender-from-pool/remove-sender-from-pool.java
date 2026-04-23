package com.example.usage;

import com.twilio.comms.TwilioCommsClient;

public class Example {
    public static void main(String[] args) {
        TwilioCommsClient client = TwilioCommsClient
            .builder()
            .credentials("<username>", "<password>")
            .build();

        client.senderPools().removeSender("comms_senderpool_01h9krwprkeee8fzqspvwy6nq8", "comms_sender_01h9krwprkeee8fzqspvwy6nq8");
    }
}