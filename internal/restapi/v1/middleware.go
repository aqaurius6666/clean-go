package v1

import (
	"strings"

	"github.com/aqaurius6666/clean-go/internal/usecases"
	"github.com/aqaurius6666/clean-go/internal/var/e"
	"github.com/aqaurius6666/clean-go/pkg/jwt"
	"github.com/aqaurius6666/clean-go/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Middleware struct {
	L        *logrus.Logger
	Usecases usecases.Usecases
}

func (s *Middleware) Token(c *gin.Context) {
	ctx := c.Request.Context()
	authStr := c.GetHeader("Authorization")
	if !strings.Contains(authStr, "Bearer ") {
		response.Response401(c, errors.New(e.ErrInvalidToken))
		return
	}

	id, err := s.Usecases.VerifyToken(ctx, authStr[7:], jwt.AccessTokenType)
	if err != nil {
		response.Response401(c, err)

		return
	}
	c.Set("id", id)
	c.Next()
}

func (s *Middleware) Logger(c *gin.Context) {
	c.Next()
	path := c.Request.URL.Path
	ierr, ok := c.Get("error")
	if !ok {
		s.L.WithField("path", path).Info("success")
		return
	}
	err, ok := ierr.(error)
	if !ok {
		s.L.Errorf("%+v", err)
		return
	}
	s.L.WithError(err).Errorf("%+v", err)
}
