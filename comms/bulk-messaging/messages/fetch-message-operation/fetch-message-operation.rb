require "twiliocomms"

client = TwilioComms::Client.new(
  account_id: "TWILIO_ACCOUNT_SID",
  auth_token: "TWILIO_AUTH_TOKEN"
)

client.messages.fetch_operation(operation_id: "comms_operation_01h9krwprkeee8fzqspvwy6nq8")
