package services

import (
	"context"
	"strconv"

	"github.com/shun198/gin-crm/prisma/db"
)

func GetAllUsers(client *db.PrismaClient) ([]db.UserModel, error) {
	users, err := client.User.FindMany().Exec(context.Background())
	return users, err
}

func GetUniqueUser(userID string, client *db.PrismaClient) (*db.UserModel, error) {
	var user_id int
	user_id, _ = strconv.Atoi(userID)
	user, err := client.User.FindUnique(
		db.User.ID.Equals(user_id),
	).Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return user, err
}

func ChangeUserDetails(client *db.PrismaClient) ([]db.UserModel, error) {
	users, err := client.User.FindMany().Exec(context.Background())
	return users, err
}

func ToggleUserActive(user *db.UserModel, client *db.PrismaClient) (*db.UserModel, error) {
	user, err := client.User.FindUnique(
		db.User.ID.Equals(user.ID),
	).Update(
		db.User.IsActive.Set(!user.IsActive),
	).Exec(context.Background())
	return user, err
}
