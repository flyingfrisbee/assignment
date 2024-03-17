package movies

import (
	"github.com/gin-gonic/gin"
)

func StartRouter(e *gin.Engine) {
	r := e.Group("/movies")
}
