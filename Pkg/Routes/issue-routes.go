package routes

import (
	controllers "github.com/aryan/go-issue-backend/Pkg/Controllers"
	"github.com/gorilla/mux"
)

var RegisterIssueRoutes = func(router *mux.Router) {
	router.HandleFunc("/issue/", controllers.CreateIssue).Methods("POST")
	router.HandleFunc("/issue/", controllers.GetIssue).Methods("GET")
	router.HandleFunc("/issueFetch/{complaintId}", controllers.GetIssueById).Methods("GET")
	router.HandleFunc("/issues/{enro}", controllers.GetIssueByEnro).Methods("GET")
	router.HandleFunc("/issueUpdate/{issueId}", controllers.UpdateIssue).Methods("PUT")
}
