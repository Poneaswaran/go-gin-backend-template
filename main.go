package main

import (
	"go-backend/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	r := gin.Default()

	r.SetTrustedProxies(nil)

	// Serve favicon
	r.StaticFile("/favicon.ico", "./static/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Gin server running 🚀",
		})
	})

	r.Run(":8000")
}
