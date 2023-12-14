package main

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to Simple Gin Server")
	server := gin.Default()
	server.GET("/home", getHome)
	server.POST("/home", postHome)
	server.GET("/queryparams", getQueryParams)
	server.GET("/urldata/:name/:age", getUrlData)
	server.Run()
}

func getHome(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Welcome to Simple Gin Server Home Page."})
}

func postHome(ctx *gin.Context) {
	body := ctx.Request.Body
	value, _ := io.ReadAll(body)
	ctx.JSON(200, gin.H{"message": "Thanks for using post home API.", "Body": string(value)})
}

func getQueryParams(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Thanks for using query params API,", "name": ctx.Query("name"), "age": ctx.Query("age")})
}

func getUrlData(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Thanks for using get url data API", "name": ctx.Param("name"), "age": ctx.Param("age")})
}
