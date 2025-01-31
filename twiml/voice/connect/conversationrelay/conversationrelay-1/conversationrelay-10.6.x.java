import com.twilio.twiml.voice.Connect;
import com.twilio.twiml.voice.ConversationRelay;
import com.twilio.twiml.voice.Language;
import com.twilio.twiml.VoiceResponse;
import com.twilio.twiml.TwiMLException;


public class Example {
    public static void main(String[] args) {
        Language language = new Language.Builder().code("sv-SE").ttsProvider("amazon").voice("Elin-Neural").transcriptionProvider("google").speechModel("long").build();
        Language language2 = new Language.Builder().code("en-US").ttsProvider("google").voice("en-US-Journey-O").build();
        ConversationRelay conversationrelay = new ConversationRelay.Builder().url("wss://mywebsocketserver.com/websocket").language(language).language(language2).build();
        Connect connect = new Connect.Builder().conversationRelay(conversationrelay).build();
        VoiceResponse response = new VoiceResponse.Builder().connect(connect).build();

        try {
            System.out.println(response.toXml());
        } catch (TwiMLException e) {
            e.printStackTrace();
        }
    }
}
