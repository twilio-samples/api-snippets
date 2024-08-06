require 'twilio-ruby'

# put your own credentials here
# To set up environmental variables, see http://twil.io/secure
account_sid = ENV['TWILIO_ACCOUNT_SID']
auth_token = ENV['TWILIO_AUTH_TOKEN']

# set up a client to talk to the Twilio REST API
@client = Twilio::REST::Client.new(account_sid, auth_token)

@participant = @client.api.conferences('AgentConf12')
                      .participants
                      .create(
                        from: '+18180021216',
                        to: '+15624421212'
                      )

puts @participant.call_sid
