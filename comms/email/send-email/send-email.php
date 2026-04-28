<?php

namespace Example;

use Twilio\Comms\TwilioComms;
use Twilio\Comms\Emails\Requests\EmailsSendRequest;
use Twilio\Comms\Types\EmailAddressSender;
use Twilio\Comms\Emails\Types\EmailsSendRequestToItemAddress;
use Twilio\Comms\Types\EmailHtmlContent;
use Twilio\Comms\Types\EmailAttachmentsItem;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
);
$client->emails->send(
    new EmailsSendRequest([
        'from' => new EmailAddressSender([
            'address' => 'support@example.company.io',
            'name' => 'Cool Co Support',
        ]),
        'to' => [
            new EmailsSendRequestToItemAddress([
                'address' => 'bob@example.com',
                'name' => 'Bob Smith',
            ]),
        ],
        'content' => new EmailHtmlContent([
            'html' => '<html><body>Hey, <br/><br/><b>Cake</b></body></html>',
            'text' => 'Hey, the cake is ready.',
            'subject' => 'Re: Wedding Cake',
            'attachments' => [
                new EmailAttachmentsItem([
                    'filename' => 'filename',
                    'contentType' => 'contentType',
                    'content' => 'content',
                ]),
            ],
        ]),
    ]),
);
