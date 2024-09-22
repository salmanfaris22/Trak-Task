package main

import (
	"todo/routs"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()
	router.Use(gin.Logger())

	routs.AuthRouter(router)
	routs.UserRouter(router)

	router.GET("/api-1", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "API 1",
		})
	})
	router.Run()

}
