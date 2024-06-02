package services

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"time"

	"github.com/shun198/gin-crm/prisma/db"
)

func CheckInvitationToken(token string, client *db.PrismaClient) (*db.InvitationModel, error) {
	invitation_token, err := GetUniqueUserByInvitationToken(token, client)
	if err != nil {
		return nil, err
	}
	return invitation_token, nil
}

func CheckResetPasswordToken(token string, client *db.PrismaClient) (*db.PasswordResetModel, error) {
	reset_password_token, err := GetUniqueUserByPasswordResetToken(token, client)
	if err != nil {
		return nil, err
	}
	return reset_password_token, nil
}

func TokenGenerator(length int) (string, error) {
	b := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func GetInvitationTokenByUserID(userID int, client *db.PrismaClient) (*db.InvitationModel, error) {
	invitation_token, err := client.Invitation.FindFirst(
		db.Invitation.UserID.Equals(userID),
	).Exec(context.Background())
	// 該当するユーザが存在しないとき
	if err != nil {
		return nil, err
	}
	return invitation_token, err
}

func UpdateInvitationToken(user *db.UserModel, client *db.PrismaClient) *db.InvitationModel {
	token, err := TokenGenerator(32)
	if err != nil {
		log.Fatal(err)
	}
	invitation_token, _ := client.Invitation.CreateOne(
		db.Invitation.Token.Set(token),
		db.Invitation.Expiry.Set(time.Now().Add(24*time.Hour)),
		db.Invitation.User.Link(
			db.User.ID.Equals(user.ID),
		),
	).Exec(context.Background())
	return invitation_token
}

func CreatePasswordResetToken(user *db.UserModel, client *db.PrismaClient) *db.PasswordResetModel {
	token, err := TokenGenerator(32)
	if err != nil {
		log.Fatal(err)
	}
	passsword_reset_token, _ := client.PasswordReset.CreateOne(
		db.PasswordReset.Token.Set(token),
		db.PasswordReset.Expiry.Set(time.Now().Add(24*time.Hour)),
		db.PasswordReset.User.Link(
			db.User.ID.Equals(user.ID),
		),
	).Exec(context.Background())
	return passsword_reset_token
}
