import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "TWILIO_ACCOUNT_SID",
        authToken: "TWILIO_AUTH_TOKEN",
    });
    await client.senderPools.list({
        startDate: new Date("2024-01-15T09:30:00Z"),
        endDate: new Date("2024-01-15T09:30:00Z"),
        operationId: "comms_operation_01h9krwprkeee8fzqspvwy6nq8",
        pageToken: "pageToken",
        pageSize: 50,
    });
}
main();
