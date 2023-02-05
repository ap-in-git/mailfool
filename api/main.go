package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitializeApiRoutes() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}

	r.GET("/ping", func(c *gin.Context) {
		fmt.Printf("ClientIP: %s\n", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/test", testMail)
	err = r.Run(":5000")
	if err != nil {
		panic(err)
	}
}
