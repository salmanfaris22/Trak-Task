package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func Application() {
	mapUrl()
	router.Run()
}
