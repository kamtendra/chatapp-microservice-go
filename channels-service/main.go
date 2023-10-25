// channels-service/main.go

package main

import (
	"chat-app-microservice/channels-service/router"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Replace with your database connection string
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "12345"
		dbname   = "mychatapp"
	)
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	

	// Create a new router
	r := mux.NewRouter()

	// Set up routes using your custom router package
	router.SetupRoutes(r, db)

	// Start the server
	serverPort := ":8080"
	log.Printf("Server is running on port %s\n", serverPort)
	log.Fatal(http.ListenAndServe(serverPort, r))
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
