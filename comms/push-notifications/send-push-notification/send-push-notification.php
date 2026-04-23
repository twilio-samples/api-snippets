<?php

namespace Example;

use Twilio\TwilioComms;
use Twilio\PushNotifications\Requests\PushNotificationsSendRequest;
use Twilio\Types\PushNotificationCredentialSender;
use Twilio\Types\SendPushNotificationDirectRecipient;
use Twilio\Types\PushNotificationProvider;
use Twilio\Types\PushNotificationContentBody;
use Twilio\Types\PushNotificationPriority;

$client = new TwilioComms(
    accountId: '<username>',
    authToken: '<password>',
);
$client->pushNotifications->send(
    new PushNotificationsSendRequest([
        'from' => new PushNotificationCredentialSender([
            'fcm' => 'comms_credential_01h9krwprkeee8fzqspvwy6nq8',
            'apn' => 'comms_credential_8qn6ywvpsqzf8eeekrpwrk9h10',
        ]),
        'to' => [
            new SendPushNotificationDirectRecipient([
                'token' => 'dqWD7WEC83K41WHyufTS7:APA91bFcrVaOLqKeJfSiEutJXX2Tr9wN_tYOwYd8rFA6mYUMBFqdz9n6k3v5EpFA_ukXD89hGqG3OarzbVfdjGnLOIAQfwbQcqJkjQWMrhwElrtU1y3JLDPfnjc0eTJLxzhyYvDFopEh',
                'provider' => PushNotificationProvider::Fcm->value,
            ]),
            new SendPushNotificationDirectRecipient([
                'token' => 'utJXX2Tr9wN_tYOwYd8rFA6mYUMBFqdz9n6k3v5EpFAukXD89hGq',
                'provider' => PushNotificationProvider::Apn->value,
            ]),
        ],
        'content' => new PushNotificationContentBody([
            'title' => 'Boarding time @ TLL',
            'body' => 'Dear customer, you have 1 hour until boarding time at the Tallinn airport',
        ]),
        'priority' => PushNotificationPriority::High->value,
        'sound' => 'ring',
    ]),
);
