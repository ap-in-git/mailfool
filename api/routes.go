package api

import (
	"github.com/ap-in-git/mailfool/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initializeRoutes(r *gin.Engine, db *gorm.DB) {
	mailBoxController := controller.NewMailBoxController(db)
	authorized := r.Group("/api/v1/")
	authorized.GET("/mail-boxes", mailBoxController.Index)

}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT,PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}
