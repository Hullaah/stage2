package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context)  {
		c.Writer.Write([]byte("Hello, World"))
	})
}
 