import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "TWILIO_ACCOUNT_SID",
        authToken: "TWILIO_AUTH_TOKEN",
    });
    await client.senders.search({
        address: "+14153902337",
        channel: "SMS",
    });
}
main();
