from twilio_comms import TwilioComms

client = TwilioComms(
    account_id="TWILIO_ACCOUNT_SID",
    auth_token="TWILIO_AUTH_TOKEN",
)

client.push_notifications.credentials.create(
    credential_type="APN",
    content={
        "certificate": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tTUlJRm5UQ0NCSVdnQXdJQkFnSUlBank5SDg0OStFOHdEUVlKS29aSWh2Y05BUUVGQlFBd2daWXhDekFKQmdOVi4uLi4uQT09LS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQ==",
        "private_key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLU1JSUVwUUlCQUFLQ0FRRUF1eWYvbE5ySDljazhEbU55bzNmR2d2Q0kxbDlzK2NtQlkzV0l6K2NVRHFteGlpZVIKLi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0t"
    },
    app_name="limonade_app",
)
