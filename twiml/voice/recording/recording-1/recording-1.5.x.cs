using System;
using Twilio.TwiML;
using Twilio.TwiML.Voice;


class Example
{
    static void Main()
    {
        var response = new VoiceResponse();
        var start = new Start();
        start.Recording(channels: "dual",
            recordingStatusCallback: "https://example.com/your-callback-url");
        response.Append(start);
        response.Say("The recording has started.");

        Console.WriteLine(response.ToString());
    }
}
