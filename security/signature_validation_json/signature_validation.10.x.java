// Install the Java helper library from twilio.com/docs/java/install
import com.twilio.security.RequestValidator;

public class Example {
  // Your Auth Token from twilio.com/user/account
  public static final String AUTH_TOKEN = System.getenv("TWILIO_AUTH_TOKEN");

  public static void main(String[] args) {
    // Initialize the request validator
    RequestValidator validator = new RequestValidator(AUTH_TOKEN);

    // Store Twilio's request URL (the url of your webhook) as a variable
    // including all query parameters
    String url = "https://example.com/myapp?bodySHA256=5ccde7145dfb8f56479710896586cb9d5911809d83afbe34627818790db0aec9";

    // Store the application/json body from Twilio's request as a variable
    // In practice, this MUST include all received parameters, not a
    // hardcoded list of parameters that you receive today. New parameters
    // may be added without notice.
    String body = "{\"CallSid\":\"CA1234567890ABCDE\",\"Caller\":\"+12349013030\"}";

    // Store the X-Twilio-Signature header attached to the request as a variable
    String twilioSignature = "hqeF3G9Hrnv6/R0jOhoYDD2PPUs=";

    // Check if the incoming signature is valid for your application URL and the incoming parameters
    System.out.println(validator.validate(url, body, twilioSignature));
  }
}
