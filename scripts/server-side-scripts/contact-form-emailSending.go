package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"gopkg.in/gomail.v2"
	/**
	  I want to swap gomail.v2 for net/smtp
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

	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")

	m := gomail.NewMessage()
	m.SetHeader("From", user)
	m.SetHeader("To", "test@varmail.org")
	m.SetHeader("Subject", "Hello with gomail")
	m.SetBody("text/plain", "This is a test email with gomail.")

	d := gomail.NewDialer("mail.gmx.com", 587, user, pass)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
