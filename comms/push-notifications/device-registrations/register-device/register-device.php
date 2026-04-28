<?php

namespace Example;

use Twilio\Comms\TwilioComms;
use Twilio\Comms\PushNotifications\DeviceRegistrations\Requests\DeviceRegistrationsRegisterRequest;
use Twilio\Comms\Types\PushNotificationProvider;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
);
$client->pushNotifications->deviceRegistrations->register(
    new DeviceRegistrationsRegisterRequest([
        'userId' => 'comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8',
        'appName' => 'limonade_app',
        'token' => 'dqWD7WEC83K41WHyufTS7:APA91bFcrVaOLqKeJfSiEutJXX2Tr9wN_tYOwYd8rFA6mYUMBFqdz9n6k3v5EpFA_ukXD89hGqG3OarzbVfdjGnLOIAQfwbQcqJkjQWMrhwElrtU1y3JLDPfnjc0eTJLxzhyYvDFopEh',
        'provider' => PushNotificationProvider::Fcm->value,
    ]),
);
