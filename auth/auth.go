package auth

import (
	"fmt"
	"net/http"

	"github.com/Hullaah/stage2/handlers"
	"github.com/gin-gonic/gin"
)

func AuthMiddleWare(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		handlers.HandlerError(c, handlers.APIError{
			Status:     "Unauthorized",
			Message:    fmt.Sprintf("Cant access path: %s", c.Request.URL.Path),
			StatusCode: http.StatusUnauthorized,
		})
		return
	}
	tokenString := authHeader[len("Bearer "):]
	claims, err := ParseTokenString(tokenString)
	handlers.HandlerError(c, err)
	c.Set("userClaims", claims)
}
