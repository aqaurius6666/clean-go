package jwt

import (
	"time"

	"github.com/pkg/errors"

	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	Id   string    `json:"id"`
	Exp  int64     `json:"exp"`
	Iat  int64     `json:"iat"`
	Type TokenType `json:"type"`
}

func (s *MyClaims) Valid() error {
	if s.Id == "" {
		return errors.New("id is empty")
	}
	if s.Exp == 0 {
		return errors.New("exp is empty")
	}
	if s.Iat == 0 {
		return errors.New("iat is empty")
	}
	if s.Type == "" {
		return errors.New("type is empty")
	}
	if s.Type != AccessTokenType && s.Type != RefreshTokenType {
		return errors.New("invalid token type")
	}
	if s.Exp < time.Now().Unix() {
		return errors.New("token expired")
	}
	if s.Iat > time.Now().Unix() {
		return errors.New("token not valid yet")
	}
	return nil
}
func IssueJWT(secret, id string, tokenType TokenType, expAt time.Time) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &MyClaims{
		Id:   id,
		Exp:  expAt.Unix(),
		Iat:  now.Unix(),
		Type: tokenType,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, err
}

	func VerifyJWT(secret, tokenString string) (string, TokenType, error) {
	claims := MyClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
	if err != nil {
		return "", "", err
	}
	if err := claims.Valid(); err != nil {
		return "", "", err
	}
	return claims.Id, claims.Type, nil
}
