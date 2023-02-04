package restapi

import (
	"net/http"

	"github.com/aqaurius6666/clean-go/pkg/swagger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type Server interface {
	http.Handler
	RegisterEndpoint()
}
type RestAPIServer struct {
	G          *gin.Engine
	Logger     *logrus.Logger
	Handler    Handler
	Middleware Middleware
}

func (s *RestAPIServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.G.ServeHTTP(w, r)
}

func (s *RestAPIServer) RegisterEndpoint() {
	s.G.Use(gin.Recovery())
	s.G.Use(s.Middleware.Logger)
	s.G.Use(otelgin.Middleware("clean-go"))
	s.G.GET("/swagger/*any", swagger.SwaggerHandler("api.swagger.json"))

	authG := s.G.Group("/auth")
	authG.POST("/login", s.Handler.HandleLoginPost)
	authG.POST("/register", s.Handler.HandleRegisterPost)
	authG.POST("/refresh", s.Handler.HandleRefreshPost)

	userG := s.G.Group("/users")
	userG.GET("/me", s.Middleware.Token, s.Handler.HandleMeGet)
	userG.PUT("/me", s.Middleware.Token, s.Handler.HandleMePut)

	postG := s.G.Group("/posts")
	postG.POST("", s.Middleware.Token, s.Handler.HandlePostsPost)
	postG.GET("/me", s.Middleware.Token, s.Handler.HandlePostsMeGet)
	postG.POST("/:postId/like", s.Middleware.Token, s.Handler.HandlePostsLikePost)
	postG.POST("/:postId/dislike", s.Middleware.Token, s.Handler.HandlePostsDislikePost)

}
