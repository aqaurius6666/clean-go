package jwt

type TokenType string

var (
	AccessTokenType  TokenType = "access"
	RefreshTokenType TokenType = "refresh"
)
