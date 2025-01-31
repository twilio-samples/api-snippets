import com.twilio.twiml.voice.Connect;
import com.twilio.twiml.voice.ConversationRelay;
import com.twilio.twiml.VoiceResponse;
import com.twilio.twiml.TwiMLException;


public class Example {
    public static void main(String[] args) {
        ConversationRelay conversationrelay = new ConversationRelay.Builder().url("wss://mywebsocketserver.com/websocket").welcomeGreeting("Hi! Ask me anything!").build();
        Connect connect = new Connect.Builder().action("https://myhttpserver.com/connect_action").conversationRelay(conversationrelay).build();
        VoiceResponse response = new VoiceResponse.Builder().connect(connect).build();

        try {
            System.out.println(response.toXml());
        } catch (TwiMLException e) {
            e.printStackTrace();
        }
    }
}
