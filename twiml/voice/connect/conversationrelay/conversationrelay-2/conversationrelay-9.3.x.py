from twilio.twiml.voice_response import Connect, ConversationRelay, Parameter, VoiceResponse

response = VoiceResponse()
connect = Connect()
conversationrelay = ConversationRelay(
    url='wss://mywebsocketserver.com/websocket')
conversationrelay.parameter(name='foo', value='bar')
conversationrelay.parameter(name='hint', value='Annoyed customer')
connect.append(conversationrelay)
response.append(connect)

print(response)
