package main

import (
	"github.com/gin-gonic/gin"
)

/**
* Main application initialization
**/
func main() {
	app := gin.New()

	// Routing & Controller
	app.GET("/", indexController)
	app.GET("/hello", helloController)
	app.Run(":9000")
}
