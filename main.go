package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

var conn pgx.Conn

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte("Hello, World"))
	})
}
