// seed.go

package main

import (
	"context"
	"log"

	"github.com/shun198/gin-crm/config"
	"github.com/shun198/gin-crm/prisma/db"
)

func main() {
	// Prismaクライアントの初期化
	client := db.NewClient()

	// データベースに接続
	err := client.Prisma.Connect()
	if err != nil {
		log.Fatalf("データベースとの接続に失敗しました: %v", err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			log.Fatalf("データベースの接続の切断に失敗しました: %v", err)
		}
	}()

	password, err := config.HashPassword("test")
	if err != nil {
		log.Fatalf("パスワードのハッシュ化に失敗しました: %v", err)
	}

	// ユーザーを作成または更新
	_, err = client.User.UpsertOne(
		db.User.Email.Equals("test01@example.com"),
	).
		Create(
			db.User.Name.Set("テストユーザゼロイチ"),
			db.User.EmployeeNumber.Set("00000001"),
			db.User.Email.Set("test01@example.com"),
			db.User.Password.Set(password),
			db.User.Role.Set("ADMIN"),
			db.User.IsActive.Set(true),
			db.User.IsVerified.Set(false),
			db.User.IsSuperuser.Set(false),
		).
		Update(
			db.User.Name.Set("テストユーザゼロイチ"),
			db.User.Password.Set(password),
			db.User.Role.Set("ADMIN"),
			db.User.IsActive.Set(true),
			db.User.IsVerified.Set(false),
			db.User.IsSuperuser.Set(false),
		).
		Exec(context.TODO())
	_, err = client.User.UpsertOne(
		db.User.Email.Equals("test02@example.com"),
	).
		Create(
			db.User.Name.Set("テストユーザゼロニ"),
			db.User.EmployeeNumber.Set("00000002"),
			db.User.Email.Set("test02@example.com"),
			db.User.Password.Set(password),
			db.User.Role.Set("GENERAL"),
			db.User.IsActive.Set(true),
			db.User.IsVerified.Set(false),
			db.User.IsSuperuser.Set(false),
		).
		Update(
			db.User.Name.Set("テストユーザゼロニ"),
			db.User.Password.Set(password),
			db.User.Role.Set("GENERAL"),
			db.User.IsActive.Set(true),
			db.User.IsVerified.Set(false),
			db.User.IsSuperuser.Set(false),
		).
		Exec(context.TODO())
	if err != nil {
		log.Fatalf("ユーザーの作成に失敗しました: %v", err)
	}

	log.Println("テストデータの作成が完了しました")
}
