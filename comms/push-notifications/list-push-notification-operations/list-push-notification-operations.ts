import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "TWILIO_ACCOUNT_SID",
        authToken: "TWILIO_AUTH_TOKEN",
    });
    await client.pushNotifications.listOperations({
        pageToken: "pageToken",
        pageSize: 50,
        startDate: new Date("2024-01-15T09:30:00Z"),
        endDate: new Date("2024-01-15T09:30:00Z"),
        status: "SCHEDULED",
    });
}
main();
