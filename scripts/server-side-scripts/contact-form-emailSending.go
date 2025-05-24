package main

import (
	"bufio"
	"bytes"
	"log"
	"net/smtp"
	"os"
	"strings"
	//"gopkg.in/gomail.v2"
	/**
	  I want to swap gomail.v2 for
	  because I don't think I'll need the features
	  that are in gomail.v2
	*/)

func loadEnv(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	// this sets it to close the file before the function returns

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if kv := strings.SplitN(line, "=", 2); len(kv) == 2 {
			os.Setenv(strings.TrimSpace(kv[0]), strings.TrimSpace(kv[1]))
		}
	}
	return scanner.Err()
}

func main() {
	if err := loadEnv(".env"); err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error loading .env: %v", err)
	}

	// Load creds from env (as discussed)
	host := os.Getenv("SMTP_HOST") // e.g. "smtp.example.com"
	port := os.Getenv("SMTP_PORT") // e.g. "587"
	user := os.Getenv("SMTP_USER") // your SMTP username
	pass := os.Getenv("SMTP_PASS") // your SMTP password

	auth := smtp.PlainAuth("", user, pass, host)

	// Build the message
	var msg bytes.Buffer
	msg.WriteString("From: " + user + "\r\n")
	msg.WriteString("To: test@varmail.org\r\n")
	msg.WriteString("Subject: Hello from net/smtp!\r\n")
	msg.WriteString("MIME-Version: 1.0\r\n")
	msg.WriteString(`Content-Type: text/html; charset="UTF-8"` + "\r\n")
	msg.WriteString("\r\n") // blank line between headers and body
	msg.WriteString(`<h1>Hi there</h1><p>This is an <strong>HTML</strong> email.</p>`)

	// Send it
	addr := host + ":" + port
	if err := smtp.SendMail(addr, auth,
		user,
		[]string{"test@varmail.org"},
		msg.Bytes(),
	); err != nil {
		log.Fatalf("failed to send email: %v", err)
	}

	log.Println("Email sent successfully!")
}
