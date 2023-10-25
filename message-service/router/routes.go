// message_routes.go

package router

import (
    "github.com/gorilla/mux"
    "database/sql"
    "chat-app-microservice/message-service/handler"
)

func SetupMessageRoutes(r *mux.Router, db *sql.DB) {
    messageHandler := handler.NewMessageHandler(db)

    r.HandleFunc("/messages/text", messageHandler.CreateTextMessageHandler).Methods("POST")
    // Define routes for other message types (image, audio, video, file, etc.)
}
