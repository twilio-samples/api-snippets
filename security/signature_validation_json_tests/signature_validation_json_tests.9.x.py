# Download the twilio-python library from twilio.com/docs/python/install
from twilio.request_validator import RequestValidator
from requests.auth import HTTPDigestAuth
from requests.auth import HTTPBasicAuth
import requests
import urllib
import os

# Your Auth Token from twilio.com/user/account saved as an environment variable
# Remember never to hard code your auth token in code, browser Javascript, or distribute it in mobile apps
# To set up environmental variables, see http://twil.io/secure
auth_token = os.environ.get('TWILIO_AUTH_TOKEN')
validator = RequestValidator(auth_token)

# Replace this URL with your unique URL
url = "https://example.com/myapp?bodySHA256=5ccde7145dfb8f56479710896586cb9d5911809d83afbe34627818790db0aec9"
# User credentials if required by your web server. Change to 'HTTPBasicAuth' if needed
auth = HTTPDigestAuth('username', 'password')

params = {}
body = """{"CallSid":"CA1234567890ABCDE","Caller":"+12349013030"}"""

def test_url(method, url, params, valid):
    if valid:
        signature = validator.compute_signature(url, params)
    else:
        signature = validator.compute_signature("http://invalid.com", params)

    headers = {'X-Twilio-Signature': signature,'Content-Type: application/json'}
    response = requests.request(method, url, headers=headers, data=body, auth=auth)
    print('HTTP {0} with {1} signature returned {2}'.format(method, 'valid' if valid else 'invalid', response.status_code))

test_url('GET', url, params, True)
test_url('GET', url, params, False)
test_url('POST', url, params, True)
test_url('POST', url, params, False)