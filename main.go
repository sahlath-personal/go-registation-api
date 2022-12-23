package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("It's Aliveeee!!!")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Go registration API",
		})
	})

	r.POST("/register", func(ctx *gin.Context) {
		fullname := ctx.PostForm("fullname")
		phone := ctx.PostForm("phone")
		email := ctx.PostForm("email")

		ctx.JSON(200, gin.H{
			"status":   "success",
			"message":  "User details successfully posted",
			"fullname": fullname,
			"phone":    phone,
			"email":    email,
		})
	})

	r.Run()
}
