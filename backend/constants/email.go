package constants

type EmailUsername string

const (
	EmailUsernameAdmin        EmailUsername = "admin"
	EmailUsernameUser         EmailUsername = "user"
	EmailUsernameOwner        EmailUsername = "owner"
	EmailUsernameGuest        EmailUsername = "guest"
	EmailUsernameNotification EmailUsername = "notification"
	EmailUsernameSupport      EmailUsername = "support"
	EmailUsernameSystem       EmailUsername = "system"
	EmailUsernameInfo         EmailUsername = "info"
)
