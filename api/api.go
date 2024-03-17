package api

import (
	"assessment/api/common"
	"assessment/api/movies"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func RunAPI() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(common.UnseenPanicHandler)

	configureSubRouters(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}
}

func configureSubRouters(engine *gin.Engine) {
	movies.StartRouter(engine)
}
