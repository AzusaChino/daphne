package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.Use(func(ctx *gin.Context) {
		fmt.Printf("url is %s", ctx.Request.URL)
		ctx.Next()
	})

	app.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
    
	app.Run()
}
