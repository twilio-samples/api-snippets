using System;
using Twilio.TwiML;
using Twilio.TwiML.Voice;


class Example
{
    static void Main()
    {
        var response = new VoiceResponse();
        var start = new Start();
        start.Transcription(statusCallbackUrl: "https://example.com/your-callback-url", intelligenceService: "GAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa");
        response.Append(start);

        Console.WriteLine(response.ToString());
    }
}
