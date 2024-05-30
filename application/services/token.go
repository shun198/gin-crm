package services

import (
	"context"
	"log"
	"time"

	"github.com/shun198/gin-crm/config"
	"github.com/shun198/gin-crm/prisma/db"
)

func CreateInvitationToken(user *db.UserModel, client *db.PrismaClient) string {
	token, err := config.TokenGenerator(64)
	if err != nil {
		log.Fatal(err)
	}
	invitation, _ := client.Invitation.CreateOne(
		db.Invitation.Token.Set(token),
		db.Invitation.Expiry.Set(time.Now().Add(24*time.Hour)),
	).Exec(context.Background())
	return invitation.Token
}

func CreatePasswordResetToken(user *db.UserModel, client *db.PrismaClient) string {
	token, err := config.TokenGenerator(64)
	if err != nil {
		log.Fatal(err)
	}
	password_reset, _ := client.PasswordReset.CreateOne(
		db.PasswordReset.Token.Set(token),
		db.PasswordReset.Expiry.Set(time.Now().Add(30*time.Minute)),
	).Exec(context.Background())
	return password_reset.Token
}

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
