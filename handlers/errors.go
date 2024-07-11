package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ValidationError struct {
	Errors []struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	} `json:"errors"`
	StatusCode int `json:"-"`
}

func (verr ValidationError) Error() string {
	return "Validation Error"
}

type APIError struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func (rerr APIError) Error() string {
	return "API error"
}

func HandlerError(c *gin.Context, err error) {
	var validationError *ValidationError
	var requestError *APIError
	if ok := errors.As(err, &validationError); ok {
		c.JSON(validationError.StatusCode, validationError)
	} else if ok := errors.As(err, &requestError); ok {
		c.JSON(requestError.StatusCode, requestError)
	} else {
		c.JSON(http.StatusInternalServerError, APIError{
			Status:     "SERVER_ERROR",
			Message:    "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
		})
	}
	c.Abort()
}
