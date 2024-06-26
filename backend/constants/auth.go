package constants

type AuthMode string

const (
	AuthModeCookie AuthMode = "Cookie"
	AuthModeXToken AuthMode = "X-Token"
)
