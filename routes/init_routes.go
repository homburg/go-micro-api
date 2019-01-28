package routes

import (
	"log"
	"net/http"

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
	log.Fatal(http.ListenAndServe(":8085", r))

	// TODO:
	// Error endpoints to slack channel from byrd web app

	// scripts
	// @DEPRECATED r.GET("v1/scripts/withdrawable")

}
