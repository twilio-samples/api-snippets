import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "<username>",
        authToken: "<password>",
    });
    await client.pushNotifications.apps.list({
        startDate: new Date("2024-01-15T09:30:00Z"),
        endDate: new Date("2024-01-15T09:30:00Z"),
        isDefault: true,
        pageToken: "pageToken",
        pageSize: 50,
    });
}
main();
