<?php

namespace Example;

use Twilio\Comms\TwilioComms;
use Twilio\Comms\Messages\Requests\MessagesListRequest;
use DateTime;
use Twilio\Comms\Types\MessageSenderChannel;
use Twilio\Comms\Types\MessageStatus;

$client = new TwilioComms(
    accountId: 'TWILIO_ACCOUNT_SID',
    authToken: 'TWILIO_AUTH_TOKEN',
);
$client->messages->list(
    new MessagesListRequest([
        'operationId' => 'comms_operation_01h9krwprkeee8fzqspvwy6nq8',
        'sessionId' => 'comms_session_01h9krwprkeee8fzqspvwy6nq8',
        'startDate' => new DateTime('2024-01-15T09:30:00Z'),
        'endDate' => new DateTime('2024-01-15T09:30:00Z'),
        'profile' => 'mem_profile_01h9krwprkeee8fzqspvwy6nq8',
        'channel' => MessageSenderChannel::Sms->value,
        'status' => MessageStatus::Queued->value,
        'tags' => 'key_1:value;key_2:value;',
        'pageToken' => 'pageToken',
        'pageSize' => 50,
    ]),
);
