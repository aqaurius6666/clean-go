package restapi

import "github.com/gin-gonic/gin"

type Handler interface {
	HandleGetUser(g *gin.Context)
}
