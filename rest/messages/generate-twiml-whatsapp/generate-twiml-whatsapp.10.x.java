import com.twilio.twiml.MessagingResponse;
import com.twilio.twiml.messaging.Body;
import com.twilio.twiml.messaging.Message;

import static spark.Spark.*;

public class WhatsAppApp {
    public static void main(String[] args) {
        get("/", (req, res) -> "Hello Web");

        post("/whatsapp", (req, res) -> {
            res.type("application/xml");
            Body body = new Body
                    .Builder("Message received! Hello again from the Twilio Sandbox for WhatsApp.")
                    .build();
            Message whatsapp = new Message
                    .Builder()
                    .body(body)
                    .build();
            MessagingResponse twiml = new MessagingResponse
                    .Builder()
                    .message(whatsapp)
                    .build();
            return twiml.toXml();
        });
    }
}