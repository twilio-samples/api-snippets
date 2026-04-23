import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "<username>",
        authToken: "<password>",
    });
    await client.pushNotifications.credentials.update("comms_credential_01h9krwprkeee8fzqspvwy6nq8", {
        isDefault: true,
    });
}
main();
