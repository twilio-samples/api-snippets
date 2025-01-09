from twilio.twiml.voice_response import VoiceResponse, Start, Transcription

response = VoiceResponse()
start = Start()
start.transcription(
    statusCallbackUrl='https://example.com/your-callback-url',
    intelligenceService='GAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa')
response.append(start)

print(response)
