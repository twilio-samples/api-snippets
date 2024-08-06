import com.twilio.jwt.accesstoken.AccessToken;
import com.twilio.jwt.accesstoken.VoiceGrant;

public class TokenGenerator {
  
  public static void main(String[] args) {
    // Get your Account SID from https://twilio.com/console
    // To set up environment variables, see http://twil.io/secure
    // Required for all types of tokens
    String twilioAccountSid = System.getenv("TWILIO_ACCOUNT_SID");
    String twilioApiKey = System.getenv("TWILIO_API_KEY");
    String twilioApiSecret = System.getenv("TWILIO_API_SECRET");

    // Required for Voice
    String outgoingApplicationSid = System.getenv("TWILIO_APP_SID");
    String identity = "user";
      
    // Create Voice grant
    VoiceGrant grant = new VoiceGrant();
    grant.setOutgoingApplicationSid(outgoingApplicationSid);

    // Optional: add to allow incoming calls
    grant.setIncomingAllow(true);
    
    // Create access token
    AccessToken token = new AccessToken.Builder(
      twilioAccountSid,
      twilioApiKey,
      twilioApiSecret
    ).identity(identity).grant(grant).build();

    System.out.println(token.toJwt());
  }
}
