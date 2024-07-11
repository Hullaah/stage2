package auth

import (
	"net/http"

	"github.com/Hullaah/stage2/db"
	"github.com/Hullaah/stage2/handlers"
	"github.com/Hullaah/stage2/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

type registerReqBody struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Phone     string `json:"phone"`
}

func RegisterHandler(c *gin.Context) {
	var reqBody registerReqBody
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		handlers.HandlerError(c, handlers.APIError{
			Status:     "Bad request",
			Message:    "Registration unsuccessful",
			StatusCode: http.StatusBadRequest,
		})
		return
	}
	if err := validate.Struct(&reqBody); err != nil {
		x, _ := err.(*validator.ValidationErrors)
		errors := make([]struct {
			Field   string `json:"field"`
			Message string `json:"message"`
		}, len(*x))
		for i, v := range *x {
			errors[i].Field = v.Field()
			errors[i].Message = v.Error()
		}
		handlers.HandlerError(c, handlers.ValidationError{
			StatusCode: http.StatusUnprocessableEntity,
			Errors:     errors,
		})
		return
	}
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(reqBody.Password), bcrypt.MinCost)
	queryEngine := db.CreateQueryEngine()
	registeredUser, _ := queryEngine.CreateUser(c, models.CreateUserParams{
		FirstName: reqBody.FirstName,
		LastName: reqBody.LastName,
		Email: reqBody.Email,
		Password: string(passwordHash),
		Phone: pgtype.Text{String: reqBody.Phone, Valid: true},
	})
	c.JSON(http.StatusOK, registeredUser)
}
