<?php

namespace Example;

use Twilio\TwilioComms;
use Twilio\Emails\Requests\EmailsSendRequest;
use Twilio\Types\EmailAddressSender;
use Twilio\Emails\Types\EmailsSendRequestToItemAddress;
use Twilio\Types\EmailHtmlContent;
use Twilio\Types\EmailAttachmentsItem;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
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
            'html' => '<html><body>Hey, <br/><br/>Cake</b></body></html>',
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
