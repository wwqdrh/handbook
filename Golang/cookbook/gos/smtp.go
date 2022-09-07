package gos

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}

func SendSimpleEmail() {
	sender := "john.doe@example.com"

	to := []string{
		"roger.roe@example.com",
	}

	user := "9c1d45eaf7af5b"
	password := "ad62926fa75d0f"

	subject := "Simple HTML mail"
	body := `<p>An old <b>falcon</b> in the sky.</p>`

	request := Mail{
		Sender:  sender,
		To:      to,
		Subject: subject,
		Body:    body,
	}

	addr := "smtp.mailtrap.io:2525"
	host := "smtp.mailtrap.io"

	msg := BuildMessage(request)
	auth := smtp.PlainAuth("", user, password, host)
	err := smtp.SendMail(addr, auth, sender, to, []byte(msg))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully")
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}
