import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "<username>",
        authToken: "<password>",
    });
    await client.pushNotifications.list({
        operationId: "comms_operation_01h9krwprkeee8fzqspvwy6nq8",
        status: "SCHEDULED",
        tags: "key_1:value;key_2:value;",
        startDate: new Date("2024-01-15T09:30:00Z"),
        endDate: new Date("2024-01-15T09:30:00Z"),
        user: "comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8",
        provider: "APN",
        pageToken: "pageToken",
        pageSize: 50,
    });
}
main();
