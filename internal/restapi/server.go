package restapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server interface {
	http.Handler
	RegisterEndpoint()
}
type RestAPIServer struct {
	G           *gin.Engine
	Handler     Handler
	UserHandler UserHandler
	Middleware  Middleware
}

func (s *RestAPIServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.G.ServeHTTP(w, r)
}

func (s *RestAPIServer) RegisterEndpoint() {
	s.G.Use(gin.Recovery())
	s.G.Use(gin.Logger())
	s.G.GET("/user", s.UserHandler.Get)
}
