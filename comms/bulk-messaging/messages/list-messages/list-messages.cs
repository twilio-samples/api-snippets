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

        await client.Messages.ListAsync(
            new MessagesListRequest {
                OperationId = "comms_operation_01h9krwprkeee8fzqspvwy6nq8",
                SessionId = "comms_session_01h9krwprkeee8fzqspvwy6nq8",
                StartDate = DateTime.Parse("2024-01-15T09:30:00Z", null, DateTimeStyles.AdjustToUniversal),
                EndDate = DateTime.Parse("2024-01-15T09:30:00Z", null, DateTimeStyles.AdjustToUniversal),
                Profile = "mem_profile_01h9krwprkeee8fzqspvwy6nq8",
                Channel = MessageSenderChannel.Sms,
                Status = MessageStatus.Queued,
                Tags = "key_1:value;key_2:value;",
                PageToken = "pageToken",
                PageSize = 50
            }
        );
    }

}
