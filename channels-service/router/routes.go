package router

import (
	"chat-app-microservice/channels-service/handler"
    "database/sql"
    "github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router, db *sql.DB) *mux.Router {
    r.HandleFunc("/channels", handler.CreateChannelHandler(db)).Methods("POST")
    r.HandleFunc("/channels/{channelID:[0-9]+}", handler.CloseChannelHandler(db)).Methods("DELETE")
    r.HandleFunc("/channels", handler.QueryChannelsHandler(db)).Methods("GET")

    return r
}

