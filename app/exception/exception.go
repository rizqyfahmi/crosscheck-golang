package exception

const (
	ErrorAccessToken   = "Error access token"
	ErrorRefreshToken  = "Error refresh token"
	ErrorDatabase      = "Error database"
	ErrorStructMapping = "Error struct mapping"
	ErrorEncryption    = "Error encryption"
)

type Exception struct {
	Message string
	Causes  string
}
