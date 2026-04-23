import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "<username>",
        authToken: "<password>",
    });
    await client.pushNotifications.deviceRegistrations.list({
        userId: "comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8",
        appName: "appName",
        pageToken: "pageToken",
        pageSize: 50,
    });
}
main();
