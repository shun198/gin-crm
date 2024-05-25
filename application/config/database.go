package config

import (
	"github.com/shun198/gin-crm/prisma/db"
)

func StartDatabase() (*db.PrismaClient, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}
	return client, nil
}
