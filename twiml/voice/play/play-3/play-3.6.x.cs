using System;
using Twilio.TwiML;
using Twilio.TwiML.Voice;


class Example
{
    static void Main()
    {
        var response = new VoiceResponse();
        response.Play(digits: "www3");

        Console.WriteLine(response.ToString());
    }
}
