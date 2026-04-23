import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "<username>",
        authToken: "<password>",
    });
    await client.pushNotifications.send({
        from: {
            fcm: "comms_credential_01h9krwprkeee8fzqspvwy6nq8",
            apn: "comms_credential_8qn6ywvpsqzf8eeekrpwrk9h10",
        },
        to: [
            {
                token: "dqWD7WEC83K41WHyufTS7:APA91bFcrVaOLqKeJfSiEutJXX2Tr9wN_tYOwYd8rFA6mYUMBFqdz9n6k3v5EpFA_ukXD89hGqG3OarzbVfdjGnLOIAQfwbQcqJkjQWMrhwElrtU1y3JLDPfnjc0eTJLxzhyYvDFopEh",
                provider: "FCM",
            },
            {
                token: "utJXX2Tr9wN_tYOwYd8rFA6mYUMBFqdz9n6k3v5EpFAukXD89hGq",
                provider: "APN",
            },
        ],
        content: {
            title: "Boarding time @ TLL",
            body: "Dear customer, you have 1 hour until boarding time at the Tallinn airport",
        },
        priority: "HIGH",
        sound: "ring",
    });
}
main();
