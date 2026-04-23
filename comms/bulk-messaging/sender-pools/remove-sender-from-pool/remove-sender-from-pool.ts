import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "<username>",
        authToken: "<password>",
    });
    await client.senderPools.removeSender("comms_senderpool_01h9krwprkeee8fzqspvwy6nq8", "comms_sender_01h9krwprkeee8fzqspvwy6nq8");
}
main();
