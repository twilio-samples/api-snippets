const VoiceResponse = require('twilio').twiml.VoiceResponse;

const response = new VoiceResponse();
const start = response.start();
start.recording({
    channels: 'dual',
    recordingStatusCallback: 'https://example.com/your-callback-url'
});
response.say('The recording has started.');

console.log(response.toString());
