using Twilio.Comms;
using System.Threading.Tasks;
using Twilio.Comms.PushNotifications;
using System;
using System.Globalization;

public partial class Examples
{
    public async Task Example() {
        var client = new TwilioComms(
            accountId: "TWILIO_ACCOUNT_SID",
            authToken: "TWILIO_AUTH_TOKEN"
        );

        await client.PushNotifications.Apps.ListAsync(
            new AppsListRequest {
                StartDate = DateTime.Parse("2024-01-15T09:30:00Z", null, DateTimeStyles.AdjustToUniversal),
                EndDate = DateTime.Parse("2024-01-15T09:30:00Z", null, DateTimeStyles.AdjustToUniversal),
                IsDefault = true,
                PageToken = "pageToken",
                PageSize = 50
            }
        );
    }

}
