// Install the Java helper library from twilio.com/docs/java/install
import com.twilio.Twilio;
import com.twilio.base.ResourceSet;
import com.twilio.rest.chat.v2.service.Role;

public class Example {
  // Get your Account SID and Auth Token from https://twilio.com/console
  // To set up environment variables, see http://twil.io/secure
  public static final String ACCOUNT_SID = System.getenv("TWILIO_ACCOUNT_SID");
  public static final String AUTH_TOKEN = System.getenv("TWILIO_AUTH_TOKEN");

  public static final String SERVICE_SID = System.getenv("TWILIO_SERVICE_SID");

  public static void main(String[] args) {
    // Initialize the client
    Twilio.init(ACCOUNT_SID, AUTH_TOKEN);

    //List the roles
    ResourceSet<Role> roles = Role.reader(SERVICE_SID).read();

    for (Role role : roles) {
      System.out.println(role.getFriendlyName());
    }
  }
}
