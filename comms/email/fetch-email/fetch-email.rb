require "twilio"

client = Twilio::Client.new(
  account_id: "<username>",
  auth_token: "<password>"
)

client.emails.fetch(email_id: "comms_email_01h9krwprkeee8fzqspvwy6nq8")
