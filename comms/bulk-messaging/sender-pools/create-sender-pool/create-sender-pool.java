package com.example.usage;

import com.twilio.comms.TwilioCommsClient;
import com.twilio.comms.resources.senderpools.requests.SenderPoolsCreateRequest;
import java.util.HashMap;

public class Example {
    public static void main(String[] args) {
        TwilioCommsClient client = TwilioCommsClient
            .builder()
            .credentials("<username>", "<password>")
            .build();

        client.senderPools().create(
            SenderPoolsCreateRequest
                .builder()
                .name("Sales Leads - APAC")
                .tags(
                    new HashMap<String, String>() {{
                        put("region", "APAC");
                    }}
                )
                .build()
        );
    }
}