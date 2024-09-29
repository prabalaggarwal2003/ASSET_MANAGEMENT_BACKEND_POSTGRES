package main

import (
	"log"
	"net/http"
	"os"

	routes "github.com/aryan/go-issue-backend/Pkg/Routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterIssueRoutes(r)
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9010"
	}

	log.Fatal(http.ListenAndServe(":"+port,
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT"}),
			handlers.AllowedHeaders([]string{"Content-Type"}),
		)(r)))
}
