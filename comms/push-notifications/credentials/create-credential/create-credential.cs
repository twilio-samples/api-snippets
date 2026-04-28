using Twilio.Comms;
using System.Threading.Tasks;
using Twilio.Comms.PushNotifications;

public partial class Examples
{
    public async Task Example() {
        var client = new TwilioComms(
            accountId: "TWILIO_ACCOUNT_SID",
            authToken: "TWILIO_AUTH_TOKEN"
        );

        await client.PushNotifications.Credentials.CreateAsync(
            new CredentialsCreateRequest {
                CredentialType = PushCredentialType.Apn,
                Content = new ApnCertificatePushCredential {
                    Certificate = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tTUlJRm5UQ0NCSVdnQXdJQkFnSUlBank5SDg0OStFOHdEUVlKS29aSWh2Y05BUUVGQlFBd2daWXhDekFKQmdOVi4uLi4uQT09LS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQ==",
                    PrivateKey = "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLU1JSUVwUUlCQUFLQ0FRRUF1eWYvbE5ySDljazhEbU55bzNmR2d2Q0kxbDlzK2NtQlkzV0l6K2NVRHFteGlpZVIKLi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0t"
                },
                AppName = "limonade_app"
            }
        );
    }

}
