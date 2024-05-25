package database

import (
	"github.com/shun198/gin-crm/prisma/db"
)

func StartDatabase() error {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return err
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	return nil
}
