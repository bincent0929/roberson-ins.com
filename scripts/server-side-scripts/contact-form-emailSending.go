// scripts/server-side-scripts/main.go
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"
)

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

// sendHandler processes form submissions at /send
func sendHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Extract form values
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	email := r.FormValue("email")
	subject := r.FormValue("subject")
	message := r.FormValue("message")

	if err := loadEnv(".env"); err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error loading .env: %v", err)
	}

	// change to who's receiving the email
	receiver_email := "test@varmail.org"
	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")

	// Build the email body
	var body bytes.Buffer
	body.WriteString("From: \"Customer Mailer\" <" + user + ">\r\n")
	body.WriteString("To: " + receiver_email + "\r\n")
	body.WriteString("Subject: " + subject + "\r\n")
	body.WriteString("MIME-Version: 1.0\r\n")
	body.WriteString(`Content-Type: text/plain; charset="UTF-8"` + "\r\n\r\n")
	body.WriteString(fmt.Sprintf("Name: %s %s\nEmail: %s\n\n%s", fname, lname, email, message))

	// Send via SMTP
	auth := smtp.PlainAuth("", user, pass, host)
	addr := host + ":" + port
	if err := smtp.SendMail(addr, auth, user, []string{receiver_email}, body.Bytes()); err != nil {
		log.Printf("SMTP error: %v", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Thank you, %s! Your message has been sent.", fname)
}

func main() {
	// Register the /send endpoint
	http.HandleFunc("/send", sendHandler)
	log.Println("Server starting on http://localhost:8080")

	// Listen on port 8080 for incoming HTTP requests
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
