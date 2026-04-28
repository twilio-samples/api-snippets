import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "TWILIO_ACCOUNT_SID",
        authToken: "TWILIO_AUTH_TOKEN",
    });
    await client.pushNotifications.users.list({
        pageToken: "pageToken",
        pageSize: 50,
    });
}
main();
