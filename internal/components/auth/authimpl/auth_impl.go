package authimpl

import (
	"context"
	"time"

	"github.com/aqaurius6666/clean-go/internal/components/auth"
	"github.com/aqaurius6666/clean-go/internal/components/user"
	"github.com/aqaurius6666/clean-go/internal/config"
	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/internal/var/e"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
	"github.com/aqaurius6666/clean-go/pkg/jwt"
)

type UseCaseImpl struct {
	AuthRepo   auth.Repository
	AuthConfig config.AuthConfig
	User       user.UseCase
}

func (s *UseCaseImpl) RegisterNewUser(ctx context.Context, email string, password string, name string) (string, error) {
	u, err := s.AuthRepo.InsertUser(ctx, gentity.WithExtend(&entities.User{
		Email:    email,
		Password: password,
		Name:     name,
	}, nil))
	if err != nil {
		return "", e.ErrEmailExisted
	}
	return u.ID, nil
}

func (s *UseCaseImpl) VerifyUserCredential(ctx context.Context, email string, password string) (string, error) {
	u, err := s.AuthRepo.SelectUser(ctx, gentity.WithExtend(&entities.User{
		Email:    email,
		Password: password,
	}, nil))
	if err != nil {
		return "", e.ErrCredentialWrong
	}
	return u.ID, nil
}

func (s *UseCaseImpl) IssueToken(ctx context.Context, id string) (accessToken string, refreshToken string, expAt int64, err error) {
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

func (s *UseCaseImpl) VerifyToken(ctx context.Context, token string, tokenType jwt.TokenType) (string, error) {
	id, _tokenType, err := jwt.VerifyJWT(s.AuthConfig.Secret, token)
	if err != nil {
		return "", err
	}
	if _tokenType != tokenType {
		return "", e.ErrInvalidToken
	}
	return id, nil
}
