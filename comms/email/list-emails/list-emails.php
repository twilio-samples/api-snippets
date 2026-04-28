<?php

namespace Example;

use Twilio\Comms\TwilioComms;
use Twilio\Comms\Emails\Requests\EmailsListRequest;
use DateTime;
use Twilio\Comms\Types\EmailStatus;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
);
$client->emails->list(
    new EmailsListRequest([
        'operationId' => 'comms_operation_01h9krwprkeee8fzqspvwy6nq8',
        'startDate' => new DateTime('2024-01-15T09:30:00Z'),
        'endDate' => new DateTime('2024-01-15T09:30:00Z'),
        'status' => EmailStatus::Scheduled->value,
        'tags' => 'key_1:value;key_2:value;',
        'pageToken' => 'pageToken',
        'pageSize' => 50,
    ]),
);
