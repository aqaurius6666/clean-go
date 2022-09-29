package v1

import (
	"github.com/gin-gonic/gin"
)

func (s *Handler) HandleGetUser(g *gin.Context) {
	g.JSON(200, gin.H{
		"message": "pong",
	})
}
