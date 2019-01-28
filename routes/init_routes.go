package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// InitRoutes init routes
func InitRoutes() {
	r := mux.NewRouter()
	// sendgrid mail
	r.HandleFunc("/v1/mail/send", MailHandler).Methods("GET")
	// slack
	r.HandleFunc("/v1/slack/users", GetSlackUsers).Methods("GET")
	r.HandleFunc("/v1/slack/assignment", SlackAssignmentHandler).Methods("POST")
	err := http.ListenAndServe(":8085", r)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Now listening to env: %s on port: 8085", os.Getenv("ENV"))
	// TODO:
	// Error endpoints to slack channel from byrd web app

	// scripts
	// @DEPRECATED r.GET("v1/scripts/withdrawable")

}
