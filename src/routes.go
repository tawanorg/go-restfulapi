package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Method GET
// Resouce http://host:port
func indexController(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Message: "API Version 0.1",
	})
}

// Method GET
// Resouce http://host:port/hello
func helloController(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Message: "Hello World!",
	})
}
