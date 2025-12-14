package main

import (
	"bytes"
	"fmt"
	"github.com/jordan-wright/email"
	"html/template"
	"net/smtp"
)

func parseTemplate(filepath string, data interface{}) (string, error) {
	t, err := template.ParseFiles(filepath)
	if err != nil {
		return "", err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func main() {
	e := email.NewEmail()
	e.From = "Toàn Lê <leductoan3082004@gmail.com>"
	e.To = []string{"teddylethal@gmail.com"}
	e.Subject = "Your subject"
	html, err := parseTemplate("example/mail/test.html", map[string]string{
		"ahihi": "chothanh",
	})
	if err != nil {
		panic(err)
	}
	e.HTML = []byte(html)
	fmt.Println(e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "leductoan3082004@gmail.com", "pyvufdtbjociioiz", "smtp.gmail.com")))
}
