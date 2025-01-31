from twilio.twiml.voice_response import Connect, ConversationRelay, VoiceResponse

response = VoiceResponse()
connect = Connect(action='https://myhttpserver.com/connect_action')
connect.conversation_relay(
    url='wss://mywebsocketserver.com/websocket',
    welcome_greeting='Hi! Ask me anything!')
response.append(connect)

print(response)
