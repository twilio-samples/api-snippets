# Get twilio-ruby from twilio.com/docs/ruby/install
require 'rubygems' # This line not needed for ruby > 1.8
require 'twilio-ruby'
require 'httparty'

# Get your Auth Token from https://www.twilio.com/console
# To set up environmental variables, see http://twil.io/secure
auth_token = ENV['TWILIO_AUTH_TOKEN']
@validator = Twilio::Security::RequestValidator.new(auth_token)

url = 'https://example.com/myapp?bodySHA256=5ccde7145dfb8f56479710896586cb9d5911809d83afbe34627818790db0aec9'
body = '{"CallSid":"CA1234567890ABCDE","Caller":"+12349013030"}'

def test_url(method, url, body, valid)
    signature = @validator.build_signature_for(valid ? url : "http://invalid.com", {})
    headers = {'X-Twilio-Signature': signature, 'Content-Type': 'application/json'}
    if method == 'GET'
      response = HTTParty.get(url, body: body, headers: headers)
    else
      response = HTTParty.post(url, body: body, headers: headers)
    end
    valid_str = valid ? 'valid' : 'invalid'
    puts "HTTP #{method} with #{valid_str} signature returned #{response.code}"
end

test_url('GET', url, body, true)
test_url('GET', url, body, false)
test_url('POST', url, body, true)
test_url('POST', url, body, false)