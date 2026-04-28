require "twiliocomms"

client = TwilioComms::Client.new(
  account_id: "TWILIO_ACCOUNT_SID",
  auth_token: "TWILIO_AUTH_TOKEN"
)

client.emails.delete(email_id: "comms_email_01h9krwprkeee8fzqspvwy6nq8")
