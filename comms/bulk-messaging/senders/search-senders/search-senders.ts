import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "<username>",
        authToken: "<password>",
    });
    await client.senders.search({
        address: "+14153902337",
        channel: "SMS",
    });
}
main();
