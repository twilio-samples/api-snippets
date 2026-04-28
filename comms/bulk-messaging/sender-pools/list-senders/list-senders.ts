import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "TWILIO_ACCOUNT_SID",
        authToken: "TWILIO_AUTH_TOKEN",
    });
    await client.senderPools.listSenders("comms_senderpool_01h9krwprkeee8fzqspvwy6nq8", {
        channel: "SMS",
        status: "ACTIVATED",
        startDate: new Date("2024-01-15T09:30:00Z"),
        endDate: new Date("2024-01-15T09:30:00Z"),
        pageToken: "pageToken",
        pageSize: 50,
    });
}
main();
