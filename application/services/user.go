package services

import (
	"context"

	"github.com/shun198/gin-crm/prisma/db"
)

func GetAllUsers(client *db.PrismaClient) ([]db.UserModel, error) {
	users, err := client.User.FindMany().Exec(context.Background())
	return users, err
}
