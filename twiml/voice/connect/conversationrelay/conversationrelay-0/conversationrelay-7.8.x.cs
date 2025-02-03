using System;
using Twilio.TwiML;
using Twilio.TwiML.Voice;


class Example
{
    static void Main()
    {
        var response = new VoiceResponse();
        var connect = new Connect(action: new Uri("https://myhttpserver.com/connect_action"));
        connect.ConversationRelay(url: "wss://mywebsocketserver.com/websocket", welcomeGreeting: "Hi! Ask me anything!");
        response.Append(connect);

        Console.WriteLine(response.ToString());
    }
}
