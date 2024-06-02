package services

import (
	"crypto/rand"
	"encoding/base64"
	"io"

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
