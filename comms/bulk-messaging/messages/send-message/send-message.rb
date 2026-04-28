require "twiliocomms"

client = TwilioComms::Client.new(
  account_id: "TWILIO_ACCOUNT_SID",
  auth_token: "TWILIO_AUTH_TOKEN"
)

client.messages.send_(
  to: [{
    address: "+12065558844",
    channel: "PHONE"
  }],
  content: {
    text: "Hello, World!"
  },
  from: {
    address: "+14153901002",
    channel: "SMS"
  }
)
