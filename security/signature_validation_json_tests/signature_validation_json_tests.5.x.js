// Get twilio-node from twilio.com/docs/libraries/node
const webhooks = require('twilio/lib/webhooks/webhooks');
const request = require('request');

// Your Auth Token from twilio.com/console
const authToken = process.env.TWILIO_AUTH_TOKEN;

// The Twilio request URL
const url = 'https://example.com/myapp?bodySHA256=5ccde7145dfb8f56479710896586cb9d5911809d83afbe34627818790db0aec9';

// The post variables in Twilio's request
const params = {};
const body = "{\"CallSid\":\"CA1234567890ABCDE\",\"Caller\":\"+12349013030\"}";


function testUrl(method, url, params, valid) {
  const signatureUrl = valid ? url : "http://invalid.com"; 
  const signature = webhooks.getExpectedTwilioSignature(authToken, signatureUrl, params);
  const options = {
      method: method,
      url: url,
      body: body,
      headers: {
        'X-Twilio-Signature': signature,
        'Content-Type': 'application/json'
      }
  }

  request(options, function(error, response, body){
      const validStr = valid ? "valid" : "invalid";
      console.log(`HTTP ${method} with ${validStr} signature returned ${response.statusCode}`);
  });
}

testUrl('GET', url, params, true);
testUrl('GET', url, params, false);
testUrl('POST', url, params, true);
testUrl('POST', url, params, false);
