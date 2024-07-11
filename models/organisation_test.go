package models_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Hullaah/stage2/db"
	"github.com/Hullaah/stage2/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

func TestGetUserIfInSameOrganisationUsersNotInSameOrganisation(t *testing.T) {
	var queryEngine = db.CreateQueryEngine()
	ctx := context.Background()

	user1, _ := queryEngine.CreateUser(ctx,
		models.CreateUserParams{
			FirstName: "Umar",
			LastName:  "Adelowo",
			Password:  "Whatever",
			Phone:     pgtype.Text{String: "11111111111", Valid: true},
			Email:     "hullaah@samdoe.com",
		})
	org1, _ := queryEngine.CreateOrganisation(ctx,
		models.CreateOrganisationParams{
			Name:        fmt.Sprintf("%s's Organisation", user1.FirstName),
			Description: pgtype.Text{String: fmt.Sprintf("Welcome to %s's organisation.", user1.FirstName), Valid: true},
		})
	queryEngine.AddUserToOrganisation(ctx,
		models.AddUserToOrganisationParams{
			UserID: user1.UserID,
			OrgID:  org1.OrgID,
		})
	user2, _ := queryEngine.CreateUser(ctx,
		models.CreateUserParams{
			FirstName: "Umar",
			LastName:  "Adelowo",
			Password:  "Whatever",
			Phone:     pgtype.Text{String: "11111111111", Valid: true},
			Email:     "hullaah@minimo.com",
		})
	org2, _ := queryEngine.CreateOrganisation(ctx,
		models.CreateOrganisationParams{
			Name:        fmt.Sprintf("%s's Organisation", user2.FirstName),
			Description: pgtype.Text{String: fmt.Sprintf("Welcome to %s's organisation.", user2.FirstName), Valid: true},
		})
	queryEngine.AddUserToOrganisation(ctx,
		models.AddUserToOrganisationParams{
			UserID: user2.UserID,
			OrgID:  org2.OrgID,
		})

	_, err := queryEngine.GetUserIfInSameOrganisation(ctx,
		models.GetUserIfInSameOrganisationParams{
			UserID:   user1.UserID,
			UserID_2: user2.UserID,
		})
	assert.ErrorIs(t, err, pgx.ErrNoRows)
}

func TestGetUserIfInSameOrganisationUsersInSameOrganisation(t *testing.T) {
	var queryEngine = db.CreateQueryEngine()
	ctx := context.Background()

	user1, _ := queryEngine.CreateUser(ctx,
		models.CreateUserParams{
			FirstName: "Umar",
			LastName:  "Adelowo",
			Password:  "Whatever",
			Phone:     pgtype.Text{String: "11111111111", Valid: true},
			Email:     "hullaah@samdoe.com",
		})
	org, _ := queryEngine.CreateOrganisation(ctx,
		models.CreateOrganisationParams{
			Name:        fmt.Sprintf("%s's Organisation", user1.FirstName),
			Description: pgtype.Text{String: fmt.Sprintf("Welcome to %s's organisation.", user1.FirstName), Valid: true},
		})
	queryEngine.AddUserToOrganisation(ctx,
		models.AddUserToOrganisationParams{
			UserID: user1.UserID,
			OrgID:  org.OrgID,
		})
	user2, _ := queryEngine.CreateUser(ctx,
		models.CreateUserParams{
			FirstName: "Umar",
			LastName:  "Adelowo",
			Password:  "Whatever",
			Phone:     pgtype.Text{String: "11111111111", Valid: true},
			Email:     "hullaah@minimo.com",
		})
	queryEngine.AddUserToOrganisation(ctx,
		models.AddUserToOrganisationParams{
			UserID: user2.UserID,
			OrgID:  org.OrgID,
		})

	user, _ := queryEngine.GetUserIfInSameOrganisation(ctx,
		models.GetUserIfInSameOrganisationParams{
			UserID:   user1.UserID,
			UserID_2: user2.UserID,
		})
	assert.Equal(t, user, user1)
}

func TestGetUserOrganisationsGetsUserOrganisationsOnly(t *testing.T) {

}
