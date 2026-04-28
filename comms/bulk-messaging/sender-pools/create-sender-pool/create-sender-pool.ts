import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "TWILIO_ACCOUNT_SID",
        authToken: "TWILIO_AUTH_TOKEN",
    });
    await client.senderPools.create({
        name: "Sales Leads - APAC",
        tags: {
            "region": "APAC",
        },
    });
}
main();
