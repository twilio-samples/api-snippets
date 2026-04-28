require "twiliocomms"

client = TwilioComms::Client.new(
  account_id: "TWILIO_ACCOUNT_SID",
  auth_token: "TWILIO_AUTH_TOKEN"
)

client.senders.resolve(recipient_addresses: [{
  address: "+14153902337",
  channel: "PHONE"
}, {
  address: "+14153902337",
  channel: "WHATSAPP"
}, {
  address: "davidpletnjov@example.com",
  channel: "EMAIL"
}])
