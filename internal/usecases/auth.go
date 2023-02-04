package usecases

import (
	"context"
	"time"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/internal/var/e"
	"github.com/aqaurius6666/clean-go/pkg/jwt"
	"github.com/pkg/errors"
)

type AuthUsecases interface {
	RegisterNewUser(ctx context.Context, email, password, name string) (string, error)
	VerifyUserCredential(ctx context.Context, email, password string) (string, error)
	IssueToken(ctx context.Context, id string) (accessToken, refreshToken string, expAt int64, err error)
	VerifyToken(ctx context.Context, token string, tokenType jwt.TokenType) (string, error)
}

func (s *UsecasesService) RegisterNewUser(ctx context.Context, email string, password string, name string) (string, error) {
	ctx, span := s.TraceProvider.Tracer(pkgName).Start(ctx, "UsecasesService.RegisterNewUser")
	defer span.End()
	u, err := s.Repo.InsertUser(ctx, &entities.User{
		Email:    email,
		Password: password,
		Name:     name,
	})
	if err != nil {
		return "", errors.New(e.ErrEmailExisted)
	}
	return u.ID, nil
}

func (s *UsecasesService) VerifyUserCredential(ctx context.Context, email string, password string) (string, error) {
	ctx, span := s.TraceProvider.Tracer(pkgName).Start(ctx, "UsecasesService.VerifyUserCredential")
	defer span.End()
	u, err := s.Repo.GetUserByEmailAndPassword(ctx, email, password)
	if err != nil {
		return "", errors.WithMessage(err, e.ErrCredentialWrong)
	}
	return u.ID, nil
}

func (s *UsecasesService) IssueToken(ctx context.Context, id string) (accessToken string, refreshToken string, expAt int64, err error) {
	ctx, span := s.TraceProvider.Tracer(pkgName).Start(ctx, "UsecasesService.IssueToken")
	defer span.End()
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
	ctx, span := s.TraceProvider.Tracer(pkgName).Start(ctx, "UsecasesService.VerifyToken")
	defer span.End()
	id, _tokenType, err := jwt.VerifyJWT(s.AuthConfig.Secret, token)
	if err != nil {
		return "", errors.WithStack(err)
	}
	if _tokenType != tokenType {
		return "", errors.New(e.ErrInvalidToken)
	}
	return id, nil
}
