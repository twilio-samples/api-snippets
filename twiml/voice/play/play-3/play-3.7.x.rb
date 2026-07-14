require 'twilio-ruby'

response = Twilio::TwiML::VoiceResponse.new
response.play(digits: 'www3')

puts response
