package jobs

import "fmt"

const (
	EmailJob = "email"
)

func SendWelcomeEmail(name, email string) error {
	fmt.Printf("Sending welcome email to %s at %s\n", name, email)
	// In a real application, you would use an email library to send the email
	return nil
}