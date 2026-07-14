const VoiceResponse = require('twilio').twiml.VoiceResponse;

const response = new VoiceResponse();
response.play({
  digits: 'www3',
});

console.log(response.toString());
