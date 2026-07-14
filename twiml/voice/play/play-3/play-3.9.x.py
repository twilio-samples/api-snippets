from twilio.twiml.voice_response import Play, VoiceResponse

response = VoiceResponse()
response.play('', digits='www3')

print(response)
