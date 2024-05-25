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

// userIDから該当する一意のユーザを取得
//
// 該当するユーザが存在すればuserを返し、存在しなければerrorを返す
func GetUniqueUser(userID string, client *db.PrismaClient) (*db.UserModel, error) {
	var user_id int
	// 数字以外のIDを入れたとき
	user_id, err := strconv.Atoi(userID)
	if err != nil {
		return nil, err
	}
	user, err := client.User.FindUnique(
		db.User.ID.Equals(user_id),
	).Exec(context.Background())
	// 該当するユーザが存在しないとき
	if err != nil {
		return nil, err
	}
	return user, err
}

func ChangeUserDetails(user *db.UserModel, client *db.PrismaClient) (*db.UserModel, error) {
	user, err := client.User.FindUnique(
		db.User.ID.Equals(user.ID),
	).Update(
		db.User.IsActive.Set(!user.IsActive),
	).Exec(context.Background())
	return user, err
}

func ToggleUserActive(user *db.UserModel, client *db.PrismaClient) (*db.UserModel, error) {
	user, err := client.User.FindUnique(
		db.User.ID.Equals(user.ID),
	).Update(
		db.User.IsActive.Set(!user.IsActive),
	).Exec(context.Background())
	return user, err
}
