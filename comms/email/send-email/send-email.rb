require "twilio"

client = Twilio::Client.new(
  account_id: "<username>",
  auth_token: "<password>"
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
    html: "<html><body>Hey, <br/><br/>Cake</b></body></html>",
    text: "Hey, the cake is ready.",
    subject: "Re: Wedding Cake",
    attachments: [{
      filename: "filename",
      content_type: "contentType",
      content: "content"
    }]
  }
)
