package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func greet(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, fmt.Sprintf("Hello %s!", name))
}

func stranger(c *gin.Context) {
	c.String(http.StatusOK, "Hello stranger!")
}

func main() {
	router := gin.Default()

	router.GET("/hello/:name", greet)
	router.GET("/hello/", stranger)
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})
	router.Run(":8000")
}
