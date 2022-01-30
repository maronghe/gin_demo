package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Zhongger ! This is your first Gin project: go-awesome-admin",
		})
	})
	r.Run()
}
