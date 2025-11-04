require 'twilio-ruby'

response = Twilio::TwiML::VoiceResponse.new
response.start do |start|
    start.recording(channels: 'dual', recording_status_callback: 'https://example.com/your-callback-url')
end
response.say(message: 'The recording has started.')

puts response
