package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func InitializeApiRoutes(db *gorm.DB) {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(corsMiddleware())
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"ip":      c.ClientIP(),
		})
	})
	r.GET("/test", testMail)
	initializeRoutes(r, db)

	err = r.Run(":5000")
	if err != nil {
		panic(err)
	}
}
