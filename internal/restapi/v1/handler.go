package v1

import (
	"github.com/aqaurius6666/clean-go/internal/components/auth"
	"github.com/aqaurius6666/clean-go/internal/components/post"
	"github.com/aqaurius6666/clean-go/internal/components/user"
)

type Handler struct {
	Auth auth.UseCase
	User user.UseCase
	Post post.UseCase
}
