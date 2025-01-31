const VoiceResponse = require('twilio').twiml.VoiceResponse;

const response = new VoiceResponse();
const connect = response.connect();
const conversationrelay = connect.conversationRelay({
    url: 'wss://mywebsocketserver.com/websocket'
});
conversationrelay.language({
    code: 'sv-SE',
    ttsProvider: 'amazon',
    voice: 'Elin-Neural',
    transcriptionProvider: 'google',
    speechModel: 'long'
});
conversationrelay.language({
    code: 'en-US',
    ttsProvider: 'google',
    voice: 'en-US-Journey-O'
});

console.log(response.toString());
