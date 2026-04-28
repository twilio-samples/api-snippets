package com.example.usage;

import com.twilio.sdk.TwilioComms;
import com.twilio.sdk.resources.senderpools.requests.SenderPoolsUpdateRequest;
import java.util.HashMap;

public class Example {
    public static void main(String[] args) {
        TwilioComms client = TwilioComms
            .builder()
            .credentials("TWILIO_ACCOUNT_SID", "TWILIO_AUTH_TOKEN")
            .build();

        client.senderPools().update(
            "comms_senderpool_01h9krwprkeee8fzqspvwy6nq8",
            SenderPoolsUpdateRequest
                .builder()
                .name("Customer Support Senders - Tier 1")
                .tags(
                    new HashMap<String, String>() {{
                        put("tier", "1");
                    }}
                )
                .build()
        );
    }
}