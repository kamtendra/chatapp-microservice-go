// channels-service/main.go

package main

import (
	"chat-app-microservice/channels-service/router"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Replace with your database connection string
	db, err := sql.Open("postgres", "user=username dbname=mychatapp sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a new router
	r := mux.NewRouter()

	// Set up routes using your custom router package
	router.SetupRoutes(r, db)

	// Start the server
	port := ":8080"
	log.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
