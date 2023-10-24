package main

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

const (
	//	config smtp gmail
	ConfigSmtpHost   = "smtp.gmail.com"
	ConfigSmtpPort   = 587
	ConfigSenderName = "Admin Digital <topikpuding29@gmail.com>"

	//	config authentication
	ConfigAuthEmail    = "taufiqkurniawan9991@gmail.com"
	ConfigAuthPassword = "bbcy gpxn qmxj fgdn"
)

func sendMail(to []string, cc []string, subject, message string) (err error) {
	//	set body generator
	body := "From: " + ConfigSenderName + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	//	authentication setup
	auth := smtp.PlainAuth("", ConfigAuthEmail, ConfigAuthPassword, ConfigSmtpHost)

	//	generate smtp address
	smtpAddress := fmt.Sprintf("%s:%d", ConfigSmtpHost, ConfigSmtpPort)

	//	send mail process
	err = smtp.SendMail(smtpAddress, auth, ConfigAuthEmail, append(to, cc...), []byte(body))
	return err
}

func main() {
	//	set mail to
	to := []string{"topikpuding29@gmail.com", "topikguanteng29@gmail.com"}

	//	set cc
	cc := []string{"topikpuding.29@gmail.com"}

	subject := "Test Mail SMTP"
	message := "Hello this is and testing for SMTP Mail"

	//call sendMail function
	err := sendMail(to, cc, subject, message)
	if err != nil {
		return
	}
	log.Println("Success send mail to", append(to, cc...))
}
