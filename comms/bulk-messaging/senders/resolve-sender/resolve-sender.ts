import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "<username>",
        authToken: "<password>",
    });
    await client.senders.resolve({
        recipientAddresses: [
            {
                address: "+14153902337",
                channel: "PHONE",
            },
            {
                address: "+14153902337",
                channel: "WHATSAPP",
            },
            {
                address: "davidpletnjov@example.com",
                channel: "EMAIL",
            },
        ],
    });
}
main();
