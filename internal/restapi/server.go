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
	PostHandler PostHandler
	Middleware  Middleware
}

func (s *RestAPIServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.G.ServeHTTP(w, r)
}

func (s *RestAPIServer) RegisterEndpoint() {
	s.G.Use(gin.Recovery())
	s.G.Use(gin.Logger())
	s.G.GET("/user", s.UserHandler.Get)
	s.G.POST("/user", s.UserHandler.Post)
	s.G.PUT("/user/:id", s.UserHandler.Put)
	s.G.DELETE("/user/:id", s.UserHandler.Delete)

	s.G.GET("/post", s.PostHandler.Get)
	s.G.POST("/post", s.PostHandler.Post)
	s.G.PUT("/post/:id", s.PostHandler.Put)
	s.G.DELETE("/post/:id", s.PostHandler.Delete)
}
