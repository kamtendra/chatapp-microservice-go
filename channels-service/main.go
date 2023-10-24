// channels-service/main.go

package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "database/sql"
    _ "github.com/lib/pq"
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
    router.Use(commonMiddleware) // Add common middleware if needed
    router.PathPrefix("/api").Handler(http.StripPrefix("/api", router.Router))
    
    router.HandleFunc("/health", healthCheckHandler).Methods("GET")

    // Set up routes defined in router
    routerSetup := router.SetupRoutes(router, db)

    log.Fatal(http.ListenAndServe(":8080", routerSetup))
}

func commonMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Implement common middleware logic here
        next.ServeHTTP(w, r)
    })
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "OK")
}
