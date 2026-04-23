import { TwilioClient } from "twilio-comms";

async function main() {
    const client = new TwilioClient({
        accountId: "<username>",
        authToken: "<password>",
    });
    await client.emails.send({
        from: {
            address: "support@example.company.io",
            name: "Cool Co Support",
        },
        to: [
            {
                address: "bob@example.com",
                name: "Bob Smith",
            },
        ],
        content: {
            html: "<html><body>Hey, <br/><br/>Cake</b></body></html>",
            text: "Hey, the cake is ready.",
            subject: "Re: Wedding Cake",
            attachments: [
                {
                    filename: "filename",
                    contentType: "contentType",
                    content: "content",
                },
            ],
        },
    });
}
main();
