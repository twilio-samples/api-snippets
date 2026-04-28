using Twilio.Comms;
using System.Threading.Tasks;
using System;
using System.Globalization;

public partial class Examples
{
    public async Task Example() {
        var client = new TwilioComms(
            accountId: "TWILIO_ACCOUNT_SID",
            authToken: "TWILIO_AUTH_TOKEN"
        );

        await client.PushNotifications.ListOperationsAsync(
            new PushNotificationsListOperationsRequest {
                PageToken = "pageToken",
                PageSize = 50,
                StartDate = DateTime.Parse("2024-01-15T09:30:00Z", null, DateTimeStyles.AdjustToUniversal),
                EndDate = DateTime.Parse("2024-01-15T09:30:00Z", null, DateTimeStyles.AdjustToUniversal),
                Status = CommunicationOperationStatus.Scheduled
            }
        );
    }

}
