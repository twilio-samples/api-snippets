import com.twilio.twiml.voice.Recording;
import com.twilio.twiml.VoiceResponse;
import com.twilio.twiml.voice.Say;
import com.twilio.twiml.voice.Start;
import com.twilio.twiml.TwiMLException;


public class Example {
    public static void main(String[] args) {
        Recording recording = new Recording.Builder().channels("dual")
            .recordingStatusCallback("https://example.com/your-callback-url")
            .build();
        Start start = new Start.Builder().recording(recording).build();
        Say say = new Say.Builder("The recording has started.").build();
        VoiceResponse response = new VoiceResponse.Builder().start(start)
            .say(say).build();

        try {
            System.out.println(response.toXml());
        } catch (TwiMLException e) {
            e.printStackTrace();
        }
    }
}
