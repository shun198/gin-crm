package services_test

import (
	"testing"

	"github.com/shun198/gin-crm/prisma/db"
	"github.com/shun198/gin-crm/services"
	"github.com/stretchr/testify/assert"
)

func TestGetUniqueUserFromEmail(t *testing.T) {
	client, mock, ensure := db.NewMock()
	defer ensure(t)

	expected := db.UserModel{
		InnerUser: db.InnerUser{
			ID:             1,
			Name:           "テストユーザゼロイチ",
			EmployeeNumber: "00000001",
			Email:          "test01@example.com",
			Role:           db.RoleAdmin,
			IsActive:       true,
			IsVerified:     false,
			IsSuperuser:    false,
		},
	}

	mock.User.Expect(
		client.User.FindUnique(
			db.User.Email.Equals(expected.Email),
		),
	).Returns(expected)

	user, err := services.GetUniqueUserByEmail(expected.Email, client)

	assert.NoError(t, err)
	assert.Equal(t, expected, *user)
}

func TestGetUniqueUserFromEmployeeNumber(t *testing.T) {
	client, mock, ensure := db.NewMock()
	defer ensure(t)

	expected := db.UserModel{
		InnerUser: db.InnerUser{
			ID:             1,
			Name:           "テストユーザゼロイチ",
			EmployeeNumber: "00000001",
			Email:          "test01@example.com",
			Role:           db.RoleAdmin,
			IsActive:       true,
			IsVerified:     false,
			IsSuperuser:    false,
		},
	}

	mock.User.Expect(
		client.User.FindUnique(
			db.User.Email.Equals(expected.EmployeeNumber),
		),
	).Returns(expected)

	user, err := services.GetUniqueUserByEmployeeNumber(expected.EmployeeNumber, client)

	assert.NoError(t, err)
	assert.Equal(t, expected, *user)
}

func TestGetAllUsers(t *testing.T) {
	// client, mock, ensure := db.NewMock()
	client, mock, ensure := db.NewMock()
	defer ensure(t)

	// 期待するユーザー一覧データ
	expected := []db.UserModel{
		{
			InnerUser: db.InnerUser{
				ID:             1,
				Name:           "テストユーザゼロイチ",
				EmployeeNumber: "00000001",
				Email:          "test01@example.com",
				Role:           db.RoleAdmin,
				IsActive:       true,
				IsVerified:     false,
				IsSuperuser:    false,
			},
		},
		{
			InnerUser: db.InnerUser{
				ID:             2,
				Name:           "テストユーザゼロニ",
				EmployeeNumber: "00000002",
				Email:          "test02@example.com",
				Role:           db.RoleGeneral,
				IsActive:       true,
				IsVerified:     false,
				IsSuperuser:    false,
			},
		},
	}
	mock.User.Expect(
		client.User.FindMany().Omit(
			db.User.Password.Field(),
			db.User.IsSuperuser.Field(),
		),
	).ReturnsMany(expected)
	users, err := services.GetAllUsers(client)
	assert.NoError(t, err)
	assert.Equal(t, expected, users)
}

func TestConvertRolesToAdmin(t *testing.T) {
	role, err := services.ConvertRoles("管理者")
	assert.NoError(t, err)
	assert.Equal(t, role, db.RoleAdmin)
}

func TestConvertRolesToGeneral(t *testing.T) {
	role, err := services.ConvertRoles("一般")
	assert.NoError(t, err)
	assert.Equal(t, role, db.RoleGeneral)
}

func TestCannotConvertNotExistingRoles(t *testing.T) {
	_, err := services.ConvertRoles("存在しないロール")
	assert.Error(t, err)
}
