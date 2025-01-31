const VoiceResponse = require('twilio').twiml.VoiceResponse;

const response = new VoiceResponse();
const connect = response.connect();
const conversationrelay = connect.conversationRelay({
    url: 'wss://mywebsocketserver.com/websocket'
});
conversationrelay.parameter({
    name: 'foo',
    value: 'bar'
});
conversationrelay.parameter({
    name: 'hint',
    value: 'Annoyed customer'
});

console.log(response.toString());
