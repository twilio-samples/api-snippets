const express = require('express');
const { MessagingResponse } = require('twilio').twiml;

const app = express();

app.post('/whatsapp', (req, res) => {
  const twiml = new MessagingResponse();

  twiml.message('Message received! Hello again from Twilio Whatsapp.');

  res.type('text/xml').send(twiml.toString());
});

app.listen(3000, () => {
  console.log('Express server listening on port 3000');
});
