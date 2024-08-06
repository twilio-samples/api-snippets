const VoiceResponse = require('twilio').twiml.VoiceResponse;

const response = new VoiceResponse();
const connect = response.connect({
    action: 'https://example.com/yourActionUrl'
});
connect.conversation({
    serviceInstanceSid: 'ISxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx',
    statusCallback: 'https://example.com/yourStatusCallback'
});

console.log(response.toString());
