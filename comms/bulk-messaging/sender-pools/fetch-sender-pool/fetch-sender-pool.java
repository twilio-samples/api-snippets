package com.example.usage;

import com.twilio.sdk.TwilioComms;

public class Example {
    public static void main(String[] args) {
        TwilioComms client = TwilioComms
            .builder()
            .credentials("TWILIO_ACCOUNT_SID", "TWILIO_AUTH_TOKEN")
            .build();

        client.senderPools().fetch("comms_senderpool_01h9krwprkeee8fzqspvwy6nq8");
    }
}