using TwilioComms;
using System.Threading.Tasks;
using System;
using System.Globalization;

public partial class Examples
{
    public async Task Example() {
        var client = new TwilioCommsClient(
            accountId: "<username>",
            authToken: "<password>"
        );

        await client.SenderPools.ListAsync(
            new SenderPoolsListRequest {
                StartDate = DateTime.Parse("2024-01-15T09:30:00Z", null, DateTimeStyles.AdjustToUniversal),
                EndDate = DateTime.Parse("2024-01-15T09:30:00Z", null, DateTimeStyles.AdjustToUniversal),
                OperationId = "comms_operation_01h9krwprkeee8fzqspvwy6nq8",
                PageToken = "pageToken",
                PageSize = 50
            }
        );
    }

}
