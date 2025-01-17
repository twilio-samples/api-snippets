<?php
// To download and install Twilio helper library visit
// https://www.twilio.com/docs/libraries/php
require_once 'vendor/autoload.php';

use Twilio\Security\RequestValidator;
use GuzzleHttp\Client;
use GuzzleHttp\Exception\ClientException;

// Your auth token from twilio.com/user/account
// To set up environmental variables, see http://twil.io/secure
$token = getenv('TWILIO_AUTH_TOKEN');

// Initialize the validator
$validator = new RequestValidator($token);

$url = 'https://example.com/myapp?bodySHA256=5ccde7145dfb8f56479710896586cb9d5911809d83afbe34627818790db0aec9';

$body = "{\"CallSid\":\"CA1234567890ABCDE\",\"Caller\":\"+12349013030\"}";

function testUrl($method, $url, $body, $valid) {
    global $validator;
    $client = new Client();

    $signature = $validator->computeSignature($valid ? $url : 'http://invalid.com', array());
    echo $signature . "\n";
    try {
        $response = $client->request($method, $url, [
            'body' => $body,
            'headers' => [
                'X-Twilio-Signature' => $signature,
                'Content-Type' => 'application/json'
            ]
        ]);
    } catch (ClientException $e) {
        $response = $e->getResponse();
    }
    $validStr = $valid ? "valid" : "invalid";
    echo "HTTP {$method} with {$validStr} signature returned {$response->getStatusCode()}\n";
}

testUrl('GET', $url, $body, true);
testUrl('GET', $url, $body, false);
testUrl('POST', $url, $body, true);
testUrl('POST', $url, $body, false);
