// router/routes.go

package router

import (
    "github.com/gorilla/mux"
    "net/http"
    "database/sql"
    "mychatapp/model"
    "mychatapp/repository"
)

func SetupRoutes(r *mux.Router, db *sql.DB) {
    channelRepo := repository.NewChannelRepository(db)

    r.HandleFunc("/channels", createChannelHandler(channelRepo)).Methods("POST")
    r.HandleFunc("/channels/{channelID}", closeChannelHandler(channelRepo)).Methods("DELETE")
    r.HandleFunc("/channels", queryChannelsHandler(channelRepo)).Methods("GET")
}

func createChannelHandler(repo *repository.ChannelRepository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Implement the logic for creating a channel using the repository
    }
}

func closeChannelHandler(repo *repository.ChannelRepository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Implement the logic for closing a channel using the repository
    }
}

func queryChannelsHandler(repo *repository.ChannelRepository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Implement the logic for querying channels using the repository
    }
}
