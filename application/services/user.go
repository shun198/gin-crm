package services

import (
	"context"
	"strconv"

	"github.com/shun198/gin-crm/dtos"
	"github.com/shun198/gin-crm/prisma/db"
)

func CreateUser(client *db.PrismaClient) (*db.UserModel, error) {
	user, err := client.User.CreateOne(
		db.User.Name.Set("テストユーザ03"),
		db.User.EmployeeNumber.Set("00000003"),
		db.User.Email.Set("00000003"),
		db.User.Password.Set("test"),
		db.User.Role.Set("ADMIN"),
	).Exec(context.Background())
	return user, err
}

// userIDから該当する一意のユーザを取得
//
// 該当するユーザが存在すればuserを返し、存在しなければerrorを返す
func GetUniqueUserByID(userID string, client *db.PrismaClient) (*db.UserModel, error) {
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

func GetUniqueUserByEmail(email string, client *db.PrismaClient) (*db.UserModel, error) {
	user, err := client.User.FindUnique(
		db.User.Email.Equals(email),
	).Exec(context.Background())
	// 該当するユーザが存在しないとき
	if err != nil {
		return nil, err
	}
	return user, err
}

func GetUniqueUserByEmployee_number(employee_number string, client *db.PrismaClient) (*db.UserModel, error) {
	user, err := client.User.FindUnique(
		db.User.EmployeeNumber.Equals(employee_number),
	).Exec(context.Background())
	// 該当するユーザが存在しないとき
	if err != nil {
		return nil, err
	}
	return user, err
}

func GetAllUsers(client *db.PrismaClient) ([]db.UserModel, error) {
	users, err := client.User.FindMany().Omit(
		db.User.Password.Field(),
		db.User.IsSuperuser.Field(),
	).Exec(context.Background())
	return users, err
}

func ChangeUserDetails(req dtos.ChangeUserDetailsDto, userID string, client *db.PrismaClient) string {
	return "未完成"
}

func ToggleUserActive(user *db.UserModel, client *db.PrismaClient) (*db.UserModel, error) {
	user, err := client.User.FindUnique(
		db.User.ID.Equals(user.ID),
	).Update(
		db.User.IsActive.Set(!user.IsActive),
	).Exec(context.Background())
	return user, err
}

func VerifyUser() string {
	return "未完成"
}

func ChangePassword() string {
	return "未完成"
}

func ResetPassword() string {
	return "未完成"
}

func CheckInvitationToken() string {
	return "未完成"
}

func CheckResetPasswordToken() string {
	return "未完成"
}

func UserInfo() string {
	return "未完成"
}
