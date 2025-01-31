const VoiceResponse = require('twilio').twiml.VoiceResponse;

const response = new VoiceResponse();
const connect = response.connect({
    action: 'https://myhttpserver.com/connect_action'
});
connect.conversationRelay({
    url: 'wss://mywebsocketserver.com/websocket',
    welcomeGreeting: 'Hi! Ask me anything!'
});

console.log(response.toString());
