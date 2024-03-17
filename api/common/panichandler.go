package common

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UnseenPanicHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("error in api: %v\n", err)
			WriteResp(c, nil, "unexpected error occured", http.StatusInternalServerError)
		}
	}()
	c.Next()
}
