import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "TWILIO_ACCOUNT_SID",
        authToken: "TWILIO_AUTH_TOKEN",
    });
    await client.pushNotifications.fetch("comms_pushnotification_01h9krwprkeee8fzqspvwy6nq8");
}
main();
