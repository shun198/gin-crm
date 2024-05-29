package services

import (
	"context"
	"time"

	"github.com/shun198/gin-crm/prisma/db"
)

func CreateInvitationToken(user *db.UserModel, client *db.PrismaClient) string {
	invitation, _ := client.Invitation.CreateOne(
		db.Invitation.UserID.Set(user.ID),
		db.Invitation.Token.Set("test"),
		db.Invitation.Expiry.Set(time.Now().Add(24*time.Hour)),
	).Exec(context.Background())
	return invitation.Token
}

func CreatePasswordResetToken(user *db.UserModel, client *db.PrismaClient) string {
	password_reset, _ := client.PasswordReset.CreateOne(
		db.PasswordReset.UserID.Set(user.ID),
		db.PasswordReset.Token.Set("test"),
		db.PasswordReset.Expiry.Set(time.Now().Add(30*time.Minute)),
	).Exec(context.Background())
	return password_reset.Token
}
