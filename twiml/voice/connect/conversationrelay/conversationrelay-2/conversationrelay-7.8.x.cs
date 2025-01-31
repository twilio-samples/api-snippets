using System;
using Twilio.TwiML;
using Twilio.TwiML.Voice;


class Example
{
    static void Main()
    {
        var response = new VoiceResponse();
        var connect = new Connect();
        var conversationrelay = new ConversationRelay(url: "wss://mywebsocketserver.com/websocket");
        conversationrelay.Parameter(name: "foo", value: "bar");
        conversationrelay.Parameter(name: "hint", value: "Annoyed customer");
        connect.Append(conversationrelay);
        response.Append(connect);

        Console.WriteLine(response.ToString());
    }
}
