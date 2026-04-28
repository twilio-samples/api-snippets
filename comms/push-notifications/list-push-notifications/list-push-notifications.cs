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

        await client.PushNotifications.ListAsync(
            new PushNotificationsListRequest {
                OperationId = "comms_operation_01h9krwprkeee8fzqspvwy6nq8",
                Status = PushNotificationStatus.Scheduled,
                Tags = "key_1:value;key_2:value;",
                StartDate = DateTime.Parse("2024-01-15T09:30:00Z", null, DateTimeStyles.AdjustToUniversal),
                EndDate = DateTime.Parse("2024-01-15T09:30:00Z", null, DateTimeStyles.AdjustToUniversal),
                User = "comms_pushnotificationuser_01h9krwprkeee8fzqspvwy6nq8",
                Provider = PushNotificationProvider.Apn,
                PageToken = "pageToken",
                PageSize = 50
            }
        );
    }

}
