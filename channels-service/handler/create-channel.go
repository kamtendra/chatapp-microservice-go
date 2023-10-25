package handler

import (
	"encoding/json"
	"net/http"
	"database/sql"
	"chat-app-microservice/channels-service/model"
)

func CreateChannelHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var channel model.Channel
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&channel); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Insert the new channel into the database
		insertQuery := "INSERT INTO channels (created_at, is_closed) VALUES (NOW(), $1) RETURNING id, created_at, is_closed"
		err := db.QueryRow(insertQuery, false).Scan(&channel.ID, &channel.CreatedAt, &channel.IsClosed)
		if err != nil {
			http.Error(w, "Error creating channel", http.StatusInternalServerError)
			println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(channel); err != nil {
			http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		}
	}
}
