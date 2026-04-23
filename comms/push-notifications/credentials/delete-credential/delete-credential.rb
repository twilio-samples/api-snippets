require "twilio"

client = Twilio::Client.new(
  account_id: "<username>",
  auth_token: "<password>"
)

client.push_notifications.credentials.delete(credential_id: "comms_credential_01h9krwprkeee8fzqspvwy6nq8")
