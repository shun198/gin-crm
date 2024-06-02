package controllers_test

import (
	"fmt"
	"testing"

	"github.com/shun198/gin-crm/prisma/db"
	"github.com/shun198/gin-crm/services"
	"github.com/stretchr/testify/assert"
)

func TestCheckResetPasswordToken(t *testing.T) {
	// create a new mock
	// this returns a mock prisma `client` and a `mock` object to set expectations
	client, mock, ensure := db.NewMock()
	// defer calling ensure, which makes sure all of the expectations were met and actually called
	// calling this makes sure that an error is returned if there was no query happening for a given expectation
	// and makes sure that all of them succeeded
	defer ensure(t)
	fmt.Print(mock)
	// start the expectation
	// mock.Post.Expect(
	// 	// define your exact query as in your tested function
	// 	// call it with the exact arguments which you expect the function to be called with
	// 	// you can copy and paste this from your tested function, and just put specific values into the arguments
	// 	client.Post.FindUnique(
	// 	db.Post.ID.Equals("123"),
	// ),
	// ).Returns(expected)
	// sets the object which should be returned in the function call
	reset_password_token, err := services.CheckResetPasswordToken("token", client)
	if err != nil {
		t.Fatal(err)
	}
	assert := assert.New(t)
	assert.Equal(reset_password_token.IsUsed, true, "they should be equal")
}
