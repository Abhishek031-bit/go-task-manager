package workers

import (
	"log"
	"task-manager/jobs"
)

type Job struct {
	Type    string
	Payload map[string]any
}

var JobQueue chan Job

func init() {
	JobQueue = make(chan Job, 100)
	go worker()
}

func worker() {
	for job := range JobQueue {
		switch job.Type {
		case jobs.EmailJob:
			name := job.Payload["name"].(string)
			email := job.Payload["email"].(string)
			if err := jobs.SendWelcomeEmail(name, email); err != nil {
				log.Printf("Failed to send welcome email: %v", err)
			}
		}
	}
}