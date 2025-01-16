import os
# Download the twilio-python library from twilio.com/docs/python/install
from twilio.request_validator import RequestValidator

# Your Auth Token from twilio.com/user/account
auth_token = os.environ['TWILIO_AUTH_TOKEN']

# Initialize the request validator
validator = RequestValidator(auth_token)

# Store Twilio's request URL (the url of your webhook) as a variable
# including all query parameters
url = 'https://example.com/myapp?bodySHA256=5ccde7145dfb8f56479710896586cb9d5911809d83afbe34627818790db0aec9'

# Store the application/json body from Twilio's request as a variable
# In practice, this MUST include all received parameters, not a
# hardcoded list of parameters that you receive today. New parameters
# may be added without notice.
body = """{"CallSid":"CA1234567890ABCDE","Caller":"+12349013030"}"""

# Store the X-Twilio-Signature header attached to the request as a variable
twilio_signature = 'hqeF3G9Hrnv6/R0jOhoYDD2PPUs='

# Check if the incoming signature is valid for your application URL and the incoming parameters
print(validator.validate(url, body, twilio_signature))
