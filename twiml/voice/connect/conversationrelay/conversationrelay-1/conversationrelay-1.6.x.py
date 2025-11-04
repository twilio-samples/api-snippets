from twilio.twiml.voice_response import Connect, ConversationRelay, Language, VoiceResponse

response = VoiceResponse()
connect = Connect()
conversationrelay = ConversationRelay(
    url='wss://mywebsocketserver.com/websocket', language='sv-SE'
)
conversationrelay.language(
    code='sv-SE',
    tts_provider='amazon',
    voice='Elin-Neural',
    transcription_provider='google',
    speech_model='long'
)
conversationrelay.language(
    code='en-US', tts_provider='google', voice='en-US-Journey-O'
)
connect.append(conversationrelay)
response.append(connect)

print(response)
