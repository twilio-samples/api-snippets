require 'twilio-ruby'

response = Twilio::TwiML::VoiceResponse.new
response.connect do |connect|
    connect.conversation_relay(url: 'wss://mywebsocketserver.com/websocket') do |conversationrelay|
        conversationrelay.parameter(name: 'foo', value: 'bar')
        conversationrelay.parameter(name: 'hint', value: 'Annoyed customer')
end
end

puts response
