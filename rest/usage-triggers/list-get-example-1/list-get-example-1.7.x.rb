# Get twilio-ruby from twilio.com/docs/ruby/install
require 'twilio-ruby'

# Get your Account SID and Auth Token from twilio.com/console
# To set up environmental variables, see http://twil.io/secure
account_sid = ENV['TWILIO_ACCOUNT_SID']
auth_token = ENV['TWILIO_AUTH_TOKEN']
@client = Twilio::REST::Client.new(account_sid, auth_token)

# Loop over triggers and print out a property for each one
triggers = @client.usage.triggers.list(recurring: 'daily',
                                       usage_category: 'calls')

triggers.each do |trigger|
  puts trigger.current_value
end
