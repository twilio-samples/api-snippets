import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "<username>",
        authToken: "<password>",
    });
    await client.pushNotifications.users.list({
        pageToken: "pageToken",
        pageSize: 50,
    });
}
main();
