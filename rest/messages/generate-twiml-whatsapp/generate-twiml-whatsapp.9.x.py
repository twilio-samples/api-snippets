from flask import Flask, request, redirect
from twilio.twiml.messaging_response import MessagingResponse

app = Flask(__name__)

@app.route("/whatsapp", methods=['GET', 'POST'])
def whatsapp_reply():
    # Start our TwiML response
    resp = MessagingResponse()

    # Add a message
    resp.message("Message received! Hello again from the Twilio Sandbox for WhatsApp.")

    return str(resp)

if __name__ == "__main__":
    app.run(debug=True)
