package auth

import (
	"log"
	"os"
	"time"

	"github.com/Hullaah/stage2/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type UserClaims struct {
	UserID    uuid.UUID
	Email     string
	Phone     string
	FirstName string
	LastName  string
	jwt.RegisteredClaims
}

func GenerateToken(u models.User) string {
	godotenv.Load("../.env")
	claims := UserClaims{
		uuid.UUID(u.UserID.Bytes),
		u.Email,
		u.Phone.String,
		u.FirstName,
		u.LastName,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := []byte(os.Getenv("SECRET_KEY"))
	s, err := token.SignedString(key)
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func ParseToken(tokenString string) (*UserClaims, error) {
	godotenv.Load("../.env")
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		// generate an unauthorised error response
		return nil, nil // tmp
	} else if !token.Valid {
		// generate an unauthorised error response
		return nil, nil // tmp
	} else if claims, ok := token.Claims.(*UserClaims); ok {
		return claims, nil
	} else {
		// generate an unauthorised error response
		return nil, nil // tmp
	}
}
