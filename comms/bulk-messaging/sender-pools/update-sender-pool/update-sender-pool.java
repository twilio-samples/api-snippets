package com.example.usage;

import com.twilio.comms.TwilioCommsClient;
import com.twilio.comms.resources.senderpools.requests.SenderPoolsUpdateRequest;
import java.util.HashMap;

public class Example {
    public static void main(String[] args) {
        TwilioCommsClient client = TwilioCommsClient
            .builder()
            .credentials("<username>", "<password>")
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