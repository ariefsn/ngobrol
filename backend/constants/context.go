package constants

type ContextKey string

const (
	AuthEmailKey       ContextKey = "AUTH_EMAIL"
	AuthRoomCodeKey    ContextKey = "AUTH_ROOM_CODE"
	RefreshTokenCtxKey ContextKey = "REFRESH_TOKEN"
	JwtClaimsCtxKey    ContextKey = "JWT_CLAIMS"
	WriterCtxKey       ContextKey = "WRITER"
)
