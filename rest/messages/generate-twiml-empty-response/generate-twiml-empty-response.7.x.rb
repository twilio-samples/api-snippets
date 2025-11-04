require 'twilio-ruby'
require 'sinatra'

post '/sms' do
  twiml = Twilio::TwiML::MessagingResponse.new

  content_type 'text/xml'

  twiml.to_s
end
