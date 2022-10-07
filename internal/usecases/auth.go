package usecases

import (
	"context"
	"time"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/internal/var/e"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
	"github.com/aqaurius6666/clean-go/pkg/jwt"
)

type AuthUsecases interface {
	RegisterNewUser(ctx context.Context, email, password, name string) (string, error)
	VerifyUserCredential(ctx context.Context, email, password string) (string, error)
	IssueToken(ctx context.Context, id string) (accessToken, refreshToken string, expAt int64, err error)
	VerifyToken(ctx context.Context, token string, tokenType jwt.TokenType) (string, error)
}

func (s *UsecasesService) RegisterNewUser(ctx context.Context, email string, password string, name string) (string, error) {
	u, err := s.Repo.InsertUser(ctx, gentity.WithExtend(&entities.User{
		Email:    email,
		Password: password,
		Name:     name,
	}, nil))
	if err != nil {
		return "", e.ErrEmailExisted
	}
	return u.ID, nil
}

func (s *UsecasesService) VerifyUserCredential(ctx context.Context, email string, password string) (string, error) {
	u, err := s.Repo.SelectUser(ctx, gentity.WithExtend(&entities.User{
		Email:    email,
		Password: password,
	}, nil))
	if err != nil {
		return "", e.ErrCredentialWrong
	}
	return u.ID, nil
}

func (s *UsecasesService) IssueToken(ctx context.Context, id string) (accessToken string, refreshToken string, expAt int64, err error) {
	expTime := time.Now().Add(time.Duration(s.AuthConfig.ExpireDuration) * time.Second)
	access, err := jwt.IssueJWT(s.AuthConfig.Secret, id, jwt.AccessTokenType, expTime)
	if err != nil {
		return "", "", 0, err
	}
	refresh, err := jwt.IssueJWT(s.AuthConfig.Secret, id, jwt.RefreshTokenType, time.Now().Add(time.Duration(s.AuthConfig.RefreshExpireDuration)*time.Second))
	if err != nil {
		return "", "", 0, err
	}
	return access, refresh, expTime.Unix(), nil
}

func (s *UsecasesService) VerifyToken(ctx context.Context, token string, tokenType jwt.TokenType) (string, error) {
	id, _tokenType, err := jwt.VerifyJWT(s.AuthConfig.Secret, token)
	if err != nil {
		return "", err
	}
	if _tokenType != tokenType {
		return "", e.ErrInvalidToken
	}
	return id, nil
}
