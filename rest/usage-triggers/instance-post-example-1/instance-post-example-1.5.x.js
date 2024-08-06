// Download the Node helper library from twilio.com/docs/node/install
// These consts are your accountSid and authToken from https://www.twilio.com/console
// To set up environmental variables, see http://twil.io/secure
const accountSid = process.env.TWILIO_ACCOUNT_SID;
const authToken = process.env.TWILIO_AUTH_TOKEN;
const client = require('twilio')(accountSid, authToken);

client.usage
  .triggers('UT33c6aeeba34e48f38d6899ea5b765ad4')
  .update({
    friendlyName: 'Monthly Maximum Call Usage',
    callbackUrl: 'https://www.example.com/monthly-usage-trigger',
  })
  .then(trigger => console.log(trigger.dateFired));
