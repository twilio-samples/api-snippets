require "twiliocomms"

client = TwilioComms::Client.new(
  account_id: "TWILIO_ACCOUNT_SID",
  auth_token: "TWILIO_AUTH_TOKEN"
)

client.emails.send_(
  from: {
    address: "support@example.company.io",
    name: "Cool Co Support"
  },
  to: [{
    address: "bob@example.com",
    name: "Bob Smith"
  }],
  content: {
    html: "<html><body>Hey, <br/><br/><b>Cake</b></body></html>",
    text: "Hey, the cake is ready.",
    subject: "Re: Wedding Cake",
    attachments: [{
      filename: "filename",
      content_type: "contentType",
      content: "content"
    }]
  }
)
