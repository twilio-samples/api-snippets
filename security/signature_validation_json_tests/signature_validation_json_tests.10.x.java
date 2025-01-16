import org.apache.http.HttpHeaders;
import org.apache.http.HttpResponse;
import org.apache.http.entity.StringEntity;
import org.apache.http.client.methods.HttpUriRequest;
import org.apache.http.client.methods.RequestBuilder;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;

import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import javax.xml.bind.DatatypeConverter;
import java.nio.charset.StandardCharsets;
import java.util.*;

public class Application {

    // Your Auth Token from twilio.com/user/account
    public static final String AUTH_TOKEN = System.getenv("TWILIO_AUTH_TOKEN");

    // Initialize the validator
    public static void main(String[] args) throws Exception {

        // The Twilio request URL
        String url = "https://example.com/myapp?bodySHA256=5ccde7145dfb8f56479710896586cb9d5911809d83afbe34627818790db0aec9";

        // The post body in Twilio's request
        String body = "{\"CallSid\":\"CA1234567890ABCDE\",\"Caller\":\"+12349013030\"}";

        TestUrl("GET", url, body, true);
        TestUrl("GET", url, body, false);
        TestUrl("POST", url, body, true);
        TestUrl("POST", url, body, false);
    }

    private static void TestUrl(
            String method, String url, String body, boolean valid) throws Exception {
        CloseableHttpClient client = HttpClients.createDefault();

        String signature = ComputeSignature(valid ? url : "http://invalid.com");

        HttpUriRequest request;
        if("GET".equals(method)) {
            request = RequestBuilder
                .get(url)
                .setEntity(new StringEntity(body))
                .setHeader(HttpHeaders.CONTENT_TYPE, "application/json")
                .setHeader("X-Twilio-Signature", signature)
                .build();
        } else {
            request = RequestBuilder
                .post(url)
                .setEntity(new StringEntity(body))
                .setHeader(HttpHeaders.CONTENT_TYPE, "application/json")
                .setHeader("X-Twilio-Signature", signature)
                .build();
        }

        HttpResponse response = client.execute(request);

        System.out.printf("HTTP %s with %s signature returned %s\n", method, valid, response.getStatusLine());
    }

    private static String ComputeSignature(String url) throws Exception{
        StringBuilder builder = new StringBuilder(url);
        Mac mac = Mac.getInstance("HmacSHA1");
        mac.init(new SecretKeySpec(AUTH_TOKEN.getBytes(), "HmacSHA1"));
        byte[] rawHmac = mac.doFinal(builder.toString().getBytes(StandardCharsets.UTF_8));
        return DatatypeConverter.printBase64Binary(rawHmac);
    }
}
