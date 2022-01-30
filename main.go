package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"简介": "：面前的这位Girl，你是我的最爱~ biu ~ biu~ biu~ ",
		})
	})
	r.Run()
}
