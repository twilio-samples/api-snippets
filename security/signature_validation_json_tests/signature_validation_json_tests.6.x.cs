// Download the twilio-csharp library from
// https://www.twilio.com/docs/libraries/csharp#installation
using System;
using System.Collections.Generic;
using Twilio.Security;
using System.Threading.Tasks;
using System.Net.Http;
using System.Text;
using System.Security.Cryptography;

class Example
{
    // Your Auth Token from twilio.com/console
    static readonly string AuthToken = Environment.GetEnvironmentVariable("AUTH_TOKEN");

    static async Task Main(string[] args)
    {
        // The Twilio request URL
        const string url = "https://example.com/myapp?bodySHA256=5ccde7145dfb8f56479710896586cb9d5911809d83afbe34627818790db0aec9";

        const string body = "{\"CallSid\":\"CA1234567890ABCDE\",\"Caller\":\"+12349013030\"}";

        await TestUrl("POST", url, body, true);
        await TestUrl("POST", url, body, false);
        await TestUrl("GET", url, body, true);
        await TestUrl("GET", url, body, false);
    }

    private static async Task TestUrl(string method, string url, string body, bool valid) {
        var client = new HttpClient();

        var signature = ComputeSignature(valid ? url : "http://invalid.com");
        client.DefaultRequestHeaders.Add("X-Twilio-Signature", signature);
    
        var request = new HttpRequestMessage
        {
            Method = (method == "GET" ? HttpMethod.Get : HttpMethod.Post),
            RequestUri = new Uri(url),
            Content = new StringContent(
                body,
                Encoding.UTF8,
                "application/json"
            )
        };
        var response = await client.SendAsync(request);

        var validStr = valid ? "valid" : "invalid";
        Console.WriteLine($"HTTP {method} with {validStr} signature returned {response.StatusCode}");
    }

    private static string ComputeSignature(string url){
        var b = new StringBuilder(url);
        var hmac = new HMACSHA1(Encoding.UTF8.GetBytes(AuthToken));
        var hash = hmac.ComputeHash(Encoding.UTF8.GetBytes(b.ToString()));
        return Convert.ToBase64String(hash);
    }
}
