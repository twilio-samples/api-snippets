from twilio.twiml.voice_response import Recording, VoiceResponse, Say, Start

response = VoiceResponse()
start = Start()
start.recording(
    channels='dual',
    recording_status_callback='https://example.com/your-callback-url'
)
response.append(start)
response.say('The recording has started.')

print(response)
