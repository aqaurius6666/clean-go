package e

import (
	"github.com/pkg/errors"
)

var (
	ErrCredentialWrong  = errors.New("email or password wrong")
	ErrEmailExisted     = errors.New("email existed")
	ErrInvalidToken     = errors.New("invalid token")
	ErrInvalidOperation = errors.New("invalid operation")
)
