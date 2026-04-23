require "twilio"

client = Twilio::Client.new(
  account_id: "<username>",
  auth_token: "<password>"
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
