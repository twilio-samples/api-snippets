import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "TWILIO_ACCOUNT_SID",
        authToken: "TWILIO_AUTH_TOKEN",
    });
    await client.senderPools.removeSender("comms_senderpool_01h9krwprkeee8fzqspvwy6nq8", "comms_sender_01h9krwprkeee8fzqspvwy6nq8");
}
main();
