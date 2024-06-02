package services

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/shun198/gin-crm/config"
	"github.com/shun198/gin-crm/prisma/db"
	"github.com/shun198/gin-crm/serializers"
)

func CreateUser(req serializers.SendInviteUserEmailSerializer, client *db.PrismaClient) *db.InvitationModel {
	randomPassword, err := config.RandomPassword()
	if err != nil {
		log.Fatal(err)
	}
	token, err := TokenGenerator(32)
	if err != nil {
		log.Fatal(err)
	}
	role, err := ConvertRoles(*req.Role)
	if err != nil {
		log.Fatal(err)
	}
	// https://goprisma.org/docs/getting-started/advanced
	// https://goprisma.org/docs/walkthrough/transactions
	user, _ := client.User.CreateOne(
		db.User.Name.Set(*req.Name),
		db.User.EmployeeNumber.Set(*req.EmployeeNumber),
		db.User.Email.Set(*req.Email),
		db.User.Password.Set(randomPassword),
		db.User.Role.Set(role),
	).Exec(context.Background())
	invitation_token, _ := client.Invitation.CreateOne(
		db.Invitation.Token.Set(token),
		db.Invitation.Expiry.Set(time.Now().Add(24*time.Hour)),
		db.Invitation.User.Link(
			db.User.ID.Equals(user.ID),
		),
	).Exec(context.Background())
	return invitation_token
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

func GetUniqueUserByEmployeeNumber(employee_number string, client *db.PrismaClient) (*db.UserModel, error) {
	user, err := client.User.FindUnique(
		db.User.EmployeeNumber.Equals(employee_number),
	).Exec(context.Background())
	// 該当するユーザが存在しないとき
	if err != nil {
		return nil, err
	}
	return user, err
}

func GetUniqueUserByInvitationToken(token string, client *db.PrismaClient) (*db.InvitationModel, error) {
	invitation_token, err := client.Invitation.FindFirst(
		db.Invitation.Token.Equals(token),
	).Exec(context.Background())
	// 該当するユーザが存在しないとき
	if err != nil {
		return nil, err
	}
	if time.Now().After(invitation_token.Expiry) || invitation_token.IsUsed {
		return nil, errors.New("無効または有効期限切れのトークンです")
	}
	return invitation_token, err
}
func GetUniqueUserByPasswordResetToken(token string, client *db.PrismaClient) (*db.PasswordResetModel, error) {
	reset_password_token, err := client.PasswordReset.FindFirst(
		db.PasswordReset.Token.Equals(token),
	).Exec(context.Background())
	// 該当するユーザが存在しないとき
	if err != nil {
		return nil, err
	}
	if time.Now().After(reset_password_token.Expiry) || reset_password_token.IsUsed {
		return nil, errors.New("無効または有効期限切れのトークンです")
	}
	return reset_password_token, err
}

func GetAllUsers(client *db.PrismaClient) ([]db.UserModel, error) {
	users, err := client.User.FindMany().Omit(
		db.User.Password.Field(),
		db.User.IsSuperuser.Field(),
	).Exec(context.Background())
	return users, err
}

func ChangeUserDetails(req serializers.ChangeUserDetailsSerializer, userID int, client *db.PrismaClient) {
	client.User.FindUnique(
		db.User.ID.Equals(userID),
	).Update(
		db.User.Name.Set(*req.Name),
		db.User.Email.Set(*req.Email),
	).Exec(context.Background())
}

func ToggleUserActive(user *db.UserModel, client *db.PrismaClient) (*db.UserModel, error) {
	user, err := client.User.FindUnique(
		db.User.ID.Equals(user.ID),
	).Update(
		db.User.IsActive.Set(!user.IsActive),
	).Exec(context.Background())
	return user, err
}

func ConvertRoles(role string) (db.Role, error) {
	switch role {
	case "管理者":
		return db.RoleAdmin, nil
	case "一般":
		return db.RoleGeneral, nil
	default:
		return "", errors.New("存在しないロールです")
	}
}

func VerifyUser(new_password string, invitation_token *db.InvitationModel, client *db.PrismaClient) {
	newPassword, _ := config.HashPassword(new_password)
	client.Invitation.FindUnique(
		db.Invitation.ID.Equals(invitation_token.ID),
	).Update(
		db.Invitation.IsUsed.Set(true),
	).Exec(context.Background())
	client.User.FindUnique(
		db.User.ID.Equals(invitation_token.UserID),
	).Update(
		db.User.IsVerified.Set(true),
		db.User.Password.Set(newPassword),
	).Exec(context.Background())
}

func CheckPassword(user *db.UserModel, password string) bool {
	check := config.CheckPasswordHash(user.Password, password)
	return check
}

func ChangePassword(new_password string, user *db.UserModel, client *db.PrismaClient) {
	newPassword, _ := config.HashPassword(new_password)
	client.User.FindUnique(
		db.User.ID.Equals(user.ID),
	).Update(
		db.User.Password.Set(newPassword),
	).Exec(context.Background())
}

func ResetPassword(new_password string, reset_password_token *db.PasswordResetModel, client *db.PrismaClient) {
	newPassword, _ := config.HashPassword(new_password)
	client.User.FindUnique(
		db.User.ID.Equals(reset_password_token.UserID),
	).Update(
		db.User.Password.Set(newPassword),
	).Exec(context.Background())
	client.PasswordReset.FindUnique(
		db.PasswordReset.ID.Equals(reset_password_token.ID),
	).Update(
		db.PasswordReset.IsUsed.Set(true),
	).Exec(context.Background())
}
