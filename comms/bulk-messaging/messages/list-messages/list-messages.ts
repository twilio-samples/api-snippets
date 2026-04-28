import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "TWILIO_ACCOUNT_SID",
        authToken: "TWILIO_AUTH_TOKEN",
    });
    await client.messages.list({
        operationId: "comms_operation_01h9krwprkeee8fzqspvwy6nq8",
        sessionId: "comms_session_01h9krwprkeee8fzqspvwy6nq8",
        startDate: new Date("2024-01-15T09:30:00Z"),
        endDate: new Date("2024-01-15T09:30:00Z"),
        profile: "mem_profile_01h9krwprkeee8fzqspvwy6nq8",
        channel: "SMS",
        status: "QUEUED",
        tags: "key_1:value;key_2:value;",
        pageToken: "pageToken",
        pageSize: 50,
    });
}
main();
