package emails

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type UrlData struct {
	Url string
}

func SendEmail(subject string, email string, url string, templateName string) {
	wd, _ := os.Getwd()

	html, _ := template.New(fmt.Sprint(templateName, ".html")).ParseFiles(fmt.Sprint(wd, "/emails/templates/", templateName, ".html"))
	plaintext, _ := template.New(fmt.Sprint(templateName, ".txt")).ParseFiles(fmt.Sprint(wd, "/emails/templates/", templateName, ".txt"))
	var tpl bytes.Buffer
	if err := html.Execute(&tpl, UrlData{Url: url}); err != nil {
		log.Fatal(err)
	}
	htmlResult := tpl.String()
	if err := plaintext.Execute(&tpl, UrlData{Url: url}); err != nil {
		log.Fatal(err)
	}
	plaintextResult := tpl.String()
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("MAIL_ADMIN"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", htmlResult)
	m.AddAlternative("text/plain", plaintextResult)

	host := os.Getenv("MAIL_HOST")
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	username := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")

	d := gomail.NewDialer(host, port, username, password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
