const VoiceResponse = require('twilio').twiml.VoiceResponse;

const response = new VoiceResponse();
const start = response.start();
start.transcription({statusCallbackUrl: 'https://example.com/your-callback-url', intelligenceService: 'GAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa'});

console.log(response.toString());
