package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/mail.v2"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	SendEmailFromAWSSES("Hello, this is a test email", "SENDER-SAMPLE-MAIL","RECIEVER-SAMPLE-MAIL","Test Email")

}

func SendEmailFromAWSSES(messageContent,sender,reciever,subject string) {
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	if smtpUsername == "" || smtpPassword == "" {
		log.Fatalf("SMTP credentials are not set in the environment")
	}
	msg := mail.NewMessage()
	msg.SetHeader("From", sender)
	msg.SetHeader("To", reciever)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", messageContent)

	d := mail.NewDialer("email-smtp.us-east-1.amazonaws.com", 587, smtpUsername, smtpPassword)
	if err := d.DialAndSend(msg); err != nil {
		log.Fatalf("Error sending email: %v", err)
	}

	fmt.Println("Email sent successfully")
}
