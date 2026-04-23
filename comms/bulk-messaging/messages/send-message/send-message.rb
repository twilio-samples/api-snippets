require "twilio"

client = Twilio::Client.new(
  account_id: "<username>",
  auth_token: "<password>"
)

client.messages.send_(
  to: [{
    address: "+14153902337",
    channel: "PHONE"
  }],
  content: {
    text: "Hello, World!"
  }
)
