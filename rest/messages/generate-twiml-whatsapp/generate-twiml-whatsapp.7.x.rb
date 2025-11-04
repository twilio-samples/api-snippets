require 'twilio-ruby'
require 'sinatra'

post '/whatsapp' do
  twiml = Twilio::TwiML::MessagingResponse.new do |r|
    r.message body: 'Message received! Hello again from the Twilio Sandbox for WhatsApp.'
  end

  content_type 'text/xml'

  twiml.to_s
end
