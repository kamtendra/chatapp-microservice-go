// message_handler.go

package main

import (
	"chat-app-microservice/message-service/repository"
	"database/sql"
	"encoding/json"
	"net/http"
	"chat-app-microservice/message-service/model"
)

type MessageHandler struct {
    repo *repository.MessageRepository
}

func NewMessageHandler(db *sql.DB) *MessageHandler {
    return &MessageHandler{repository.NewMessageRepository(db)}
}

func (h *MessageHandler) CreateTextMessageHandler(w http.ResponseWriter, r *http.Request) {
    var message models.Message
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&message); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // Insert the text message into the database
    messageID, err := h.repo.CreateTextMessage(message.Content)
    if err != nil {
        http.Error(w, "Error creating text message", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(struct {
        MessageID int `json:"message_id"`
    }{MessageID: messageID})
}

// Implement handlers for other message types (image, audio, video, file, etc.)
