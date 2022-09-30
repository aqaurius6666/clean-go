package response

import "github.com/gin-gonic/gin"

func Response200(g *gin.Context, data interface{}) {
	g.JSON(200, gin.H{
		"success": true,
		"code":    200,
		"data":    data,
	})
}

func Response400(g *gin.Context, err error) {
	g.JSON(200, gin.H{
		"success": false,
		"code":    400,
		"error":   err.Error(),
	})
}
func Response401(g *gin.Context, err error) {
	g.JSON(200, gin.H{
		"success": false,
		"code":    401,
		"error":   err.Error(),
	})
}

func Response500(g *gin.Context, err error) {
	g.JSON(200, gin.H{
		"success": false,
		"code":    500,
		"error":   err.Error(),
	})
}
