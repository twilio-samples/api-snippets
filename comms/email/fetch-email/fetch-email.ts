import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "TWILIO_ACCOUNT_SID",
        authToken: "TWILIO_AUTH_TOKEN",
    });
    await client.emails.fetch("comms_email_01h9krwprkeee8fzqspvwy6nq8");
}
main();
