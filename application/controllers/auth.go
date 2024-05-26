package controllers

import "github.com/shun198/gin-crm/services"

func Login() {
	services.Login()
}

func Logout() {
	services.Logout()
}
