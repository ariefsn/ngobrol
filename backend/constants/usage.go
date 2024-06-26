package constants

type UsageAccountStatus string

const (
	UsageAccountStatusEmailActive          UsageAccountStatus = "active"
	UsageAccountStatusEmailInactive        UsageAccountStatus = "inactive"
	UsageAccountStatusWhatsappConnected    UsageAccountStatus = "connected"
	UsageAccountStatusWhatsappDisconnected UsageAccountStatus = "disconnected"
)
