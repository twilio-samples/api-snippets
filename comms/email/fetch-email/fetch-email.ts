import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "<username>",
        authToken: "<password>",
    });
    await client.emails.fetch("comms_email_01h9krwprkeee8fzqspvwy6nq8");
}
main();
