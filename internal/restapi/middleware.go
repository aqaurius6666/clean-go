package restapi

import "github.com/gin-gonic/gin"

type Middleware interface {
	Token(*gin.Context)
	Logger(*gin.Context)
}
