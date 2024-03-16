package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Developing Skill-MAP")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Page is under construction",
		})
	})

	r.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "login-required",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
