package com.example.usage;

import com.twilio.sdk.TwilioComms;
import com.twilio.sdk.resources.senderpools.requests.SenderPoolsCreateRequest;
import java.util.HashMap;

public class Example {
    public static void main(String[] args) {
        TwilioComms client = TwilioComms
            .builder()
            .credentials("TWILIO_ACCOUNT_SID", "TWILIO_AUTH_TOKEN")
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