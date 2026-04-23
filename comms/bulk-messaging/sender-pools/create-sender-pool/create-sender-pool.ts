import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "<username>",
        authToken: "<password>",
    });
    await client.senderPools.create({
        name: "Sales Leads - APAC",
        tags: {
            "region": "APAC",
        },
    });
}
main();
