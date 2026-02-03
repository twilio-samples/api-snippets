using System;
using Twilio.TwiML;
using Twilio.TwiML.Voice;


class Example
{
    static void Main()
    {
        var response = new VoiceResponse();
        var connect = new Connect();
        var conversationrelay = new ConversationRelay(url: "wss://mywebsocketserver.com/websocket", language: "sv-SE");
        conversationrelay.Language(code: "sv-SE", ttsProvider: "amazon",
            voice: "Elin-Neural", transcriptionProvider: "google",
            speechModel: "long");
        conversationrelay.Language(code: "en-US", ttsProvider: "google",
            voice: "en-US-Journey-O");
        connect.Append(conversationrelay);
        response.Append(connect);

        Console.WriteLine(response.ToString());
    }
}
