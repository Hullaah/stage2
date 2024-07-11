package auth_test

import (
	"time"
	"testing"

	"github.com/Hullaah/stage2/models"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"github.com/Hullaah/stage2/auth"
)

func TestTokenExpiresAtCorrectTime(t *testing.T) {
	user := models.User{
		UserID: pgtype.UUID{Bytes: uuid.New(), Valid: true},
		FirstName: "Umar",
		LastName: "Adelowo",
		Email: "umar@example.com",
		Password: "Mini  mini mani moe",
		Phone: pgtype.Text{String: "07080490736", Valid: true},
	}
	token := auth.GenerateToken(user)
	userClaim, _ := auth.ParseToken(token)
	assert.True(t, time.Now().Before(userClaim.ExpiresAt.Time))
}

func TestCorrectUserDetailsIsFoundInToken(t *testing.T)  {
	assert := assert.New(t)
	user := models.User{
		UserID: pgtype.UUID{Bytes: uuid.New(), Valid: true},
		FirstName: "Umar",
		LastName: "Adelowo",
		Email: "umar@example.com",
		Password: "Mini  mini mani moe",
		Phone: pgtype.Text{String: "07080490736", Valid: true},
	}
	token := auth.GenerateToken(user)
	userClaim, _ := auth.ParseToken(token)
	assert.Equal(user.UserID.Bytes, [16]byte(userClaim.UserID))
	assert.Equal(user.FirstName, userClaim.FirstName)
	assert.Equal(user.LastName, userClaim.LastName)
	assert.Equal(user.Email, userClaim.Email)
	assert.Equal(user.Phone.String, userClaim.Phone)

}
