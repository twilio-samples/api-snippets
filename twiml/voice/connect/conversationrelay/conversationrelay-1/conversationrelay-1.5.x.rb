require 'twilio-ruby'

response = Twilio::TwiML::VoiceResponse.new
response.connect do |connect|
    connect.conversation_relay(url: 'wss://mywebsocketserver.com/websocket', language: 'sv-SE') do |conversationrelay|
        conversationrelay.language(code: 'sv-SE', tts_provider: 'amazon', voice: 'Elin-Neural', transcription_provider: 'google', speech_model: 'long')
        conversationrelay.language(code: 'en-US', tts_provider: 'google', voice: 'en-US-Journey-O')
end
end

puts response
