package auth

import (
	"context"

	"github.com/aqaurius6666/clean-go/internal/entities"
	"github.com/aqaurius6666/clean-go/pkg/gentity"
	"github.com/aqaurius6666/clean-go/pkg/jwt"
)

type UseCase interface {
	RegisterNewUser(ctx context.Context, email, password, name string) (string, error)
	VerifyUserCredential(ctx context.Context, email, password string) (string, error)
	IssueToken(ctx context.Context, id string) (accessToken, refreshToken string, expAt int64, err error)
	VerifyToken(ctx context.Context, token string, tokenType jwt.TokenType) (string, error)
}

type Repository interface {
	SelectUser(ctx context.Context, ex gentity.Extend[*entities.User]) (*entities.User, error)
	InsertUser(ctx context.Context, ex gentity.Extend[*entities.User]) (*entities.User, error)
}
