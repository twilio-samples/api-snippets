// Download the twilio-csharp library from
// https://www.twilio.com/docs/libraries/csharp#installation
using System;
using System.Collections.Generic;
using Twilio.Security;

class Example
{
    static void Main(string[] args)
    {
        // Your Auth Token from twilio.com/console
        const string authToken = Environment.GetEnvironmentVariable("TWILIO_AUTH_TOKEN");

        // Initialize the request validator
        var validator = new RequestValidator(authToken);

        // Store Twilio's request URL (the url of your webhook) as a variable
        // including all query parameters
        const string url = "https://example.com/myapp?bodySHA256=5ccde7145dfb8f56479710896586cb9d5911809d83afbe34627818790db0aec9";

        // Store the application/json body from Twilio's request as a variable
        // In practice, this MUST include all received parameters, not a
        // hardcoded list of parameters that you receive today. New parameters
        // may be added without notice.
        const string body = "{\"CallSid\":\"CA1234567890ABCDE\",\"Caller\":\"+12349013030\"}";


        // Store the X-Twilio-Signature header attached to the request as a variable
        const string twilioSignature = "hqeF3G9Hrnv6/R0jOhoYDD2PPUs=";

        // Check if the incoming signature is valid for your application URL and the incoming body
        Console.WriteLine(validator.Validate(url, body, twilioSignature));
    }
}
