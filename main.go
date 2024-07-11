package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Hullaah/stage2/db"
)

func main() {
	queryEngine := db.CreateQueryEngine()
	r := gin.Default()
	r.GET("/", func(c *gin.Context)  {
		c.Writer.Write([]byte("Hello, World"))
	})
}
 