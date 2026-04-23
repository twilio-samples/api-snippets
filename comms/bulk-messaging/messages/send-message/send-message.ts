import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "<username>",
        authToken: "<password>",
    });
    await client.messages.send({
        to: [
            {
                address: "+14153902337",
                channel: "PHONE",
            },
        ],
        content: {
            text: "Hello, World!",
        },
    });
}
main();
