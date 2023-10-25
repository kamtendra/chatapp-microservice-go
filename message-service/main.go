// main.go

package main

import (
    "database/sql"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
    "mychatapp/router"
)

func main() {
    // Replace with your database connection string
    db, err := sql.Open("postgres", "user=username dbname=mychatapp sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    router := mux.NewRouter()

    // Set up CORS middleware
    corsHandler := cors.New(cors.Options{
        AllowedOrigins: []string{"*"}, // Replace with your allowed origins
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"Content-Type", "Authorization"},
    })

    router.Use(corsHandler.Handler)

    // Set up message routes
    routerSetup := router.SetupMessageRoutes(router, db)

    // Start the server
    port := ":8081" // You can use a different port if needed
    log.Printf("Server is running on port %s\n", port)
    log.Fatal(http.ListenAndServe(port, routerSetup))
}
