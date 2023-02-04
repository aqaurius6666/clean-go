package restapi

import "github.com/gin-gonic/gin"

type Handler interface {
	UserHandler
	AuthHandler
	PostHandler
}

type AuthHandler interface {
	HandleLoginPost(g *gin.Context)
	HandleRegisterPost(g *gin.Context)
	HandleRefreshPost(g *gin.Context)
}

type UserHandler interface {
	HandleMeGet(*gin.Context)
	HandleMePut(*gin.Context)
}

type PostHandler interface {
	HandlePostsMeGet(*gin.Context)
	HandlePostsPost(*gin.Context)
	HandlePostsLikePost(*gin.Context)
	HandlePostsDislikePost(*gin.Context)
}
