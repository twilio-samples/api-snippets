require 'twilio-ruby'

response = Twilio::TwiML::VoiceResponse.new
response.connect(action: 'https://myhttpserver.com/connect_action') do |connect|
    connect.conversation_relay(url: 'wss://mywebsocketserver.com/websocket', welcome_greeting: 'Hi! Ask me anything!')
end

puts response
