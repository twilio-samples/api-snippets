require "twilio"

client = Twilio::Client.new(
  account_id: "<username>",
  auth_token: "<password>"
)

client.senders.search(
  address: "+14153902337",
  channel: "SMS"
)
