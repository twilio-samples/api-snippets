<?php
// NOTE: This example uses the next generation Twilio helper library - for more
// information on how to download and install this version, visit
// https://www.twilio.com/docs/libraries/php
require_once '/path/to/vendor/autoload.php';

use Twilio\Security\RequestValidator;

// Your auth token from twilio.com/user/account
$token = getenv("TWILIO_AUTH_TOKEN");

// The X-Twilio-Signature header - in PHP this should be
// You may be able to use $signature = $_SERVER["HTTP_X_TWILIO_SIGNATURE"];
$signature = 'hqeF3G9Hrnv6/R0jOhoYDD2PPUs=';

// Initialize the request validator
$validator = new RequestValidator($token);

// Store Twilio's request URL (the url of your webhook) as a variable
// including all query parameters
// You may be able to use $url = $_SERVER['SCRIPT_URI']
$url = 'https://example.com/myapp?bodySHA256=5ccde7145dfb8f56479710896586cb9d5911809d83afbe34627818790db0aec9';

// Store the application/json body from Twilio's request as a variable
// In practice, this MUST include all received parameters, not a
// hardcoded list of parameters that you receive today. New parameters
// may be added without notice.
// You may be able to use $body = $_POST
$body = "{\"CallSid\":\"CA1234567890ABCDE\",\"Caller\":\"+12349013030\"}";

// Check if the incoming signature is valid for your application URL and the incoming parameters
if ($validator->validate($signature, $url, $body)) {
  echo "Confirmed to have come from Twilio.";
} else {
  echo "NOT VALID. It might have been spoofed!";
}
