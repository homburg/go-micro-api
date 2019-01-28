package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/byblix/go-micro/models"
	"github.com/byblix/go-micro/services"
)

// TODO:
// Create endpoint to send POST body from web app
// redirect the body values to slack channel and write message

// SlackAssignmentHandler to slack
func SlackAssignmentHandler(w http.ResponseWriter, r *http.Request) {
	// Get JSON from client and parse to byte array
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Host, r.URL, string(body))
	// Create http req to slack webhook
	var msg models.SlackMessage
	services.AssignmentPost(body, msg)
}

// GetSlackUsers send msg to slack about new assignment faster
func GetSlackUsers(w http.ResponseWriter, r *http.Request) {
	var slackUsers models.SlackUser
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(slackUsers)
	fmt.Println(r.Host, r.URL)
}
