import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "<username>",
        authToken: "<password>",
    });
    await client.senderPools.update("comms_senderpool_01h9krwprkeee8fzqspvwy6nq8", {
        name: "Customer Support Senders - Tier 1",
        tags: {
            "tier": "1",
        },
    });
}
main();
