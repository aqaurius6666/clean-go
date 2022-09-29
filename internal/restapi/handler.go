package restapi

import "github.com/gin-gonic/gin"

type Handler interface {
	HandleGetUser(g *gin.Context)
}

type UserHandler interface {
	Get(*gin.Context)
}
