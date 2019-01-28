package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/byblix/go-micro/models"
)

// InitBot asd
func InitBot() {
	fmt.Println(os.Getenv("SLACK_BOT"))
	// slackClient := slack.New(os.Getenv("SLACK_BOT"))
	// _ = slackClient
}

// GetSlackUsers to get all users
func GetSlackUsers() {}

// AssignmentPost asd
func AssignmentPost(body []byte, msg models.SlackMessage) {
	webHook := os.Getenv("SLACK_WEBHOOK")
	err := json.Unmarshal(body, &msg)
	req, err := http.NewRequest("POST", webHook, bytes.NewBuffer(body))
	req.Header.Set("Content-type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error %s", err)
	}
	defer resp.Body.Close()
}
