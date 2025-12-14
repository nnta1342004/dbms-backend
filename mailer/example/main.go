package main

import (
	"hareta/mailer"
	"log"
)

func main() {

	//from := "toan.le3008biot@hcmut.edu.vn"
	//password := "oktofnfqqmfhledm"
	//
	//toEmailAddress := "thanhthaofirecrush@mailer.com"
	//to := []string{toEmailAddress}
	//
	host := "smtp.gmail.com"
	port := "587"
	//address := host + ":" + port
	//
	//subject := "Subject: This is the subject of the mail\n"
	//body := "This is the body of the mail"
	//message := []byte(subject + body)
	//
	//auth := smtp.PlainAuth("", from, password, host)
	//
	//err := smtp.SendMail(address, auth, from, to, message)
	//if err != nil {
	//	panic(err)
	//}

	Sender := mailer.NewMailSender(host, port)
	data := mailer.Mail{
		Subject:  "hello toan le day",
		Body:     "test",
		Receiver: "thanhthaofirecrush@gmail.com",
	}
	if err := Sender.SendMail(data); err != nil {
		log.Fatalln(err)
	}
	log.Fatalln("success")
}
