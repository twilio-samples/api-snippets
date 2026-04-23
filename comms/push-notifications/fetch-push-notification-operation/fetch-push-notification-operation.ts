import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "<username>",
        authToken: "<password>",
    });
    await client.pushNotifications.fetchOperation("comms_operation_01h9krwprkeee8fzqspvwy6nq8");
}
main();
