// create_channel_handler.go

package main

import (
    "encoding/json"
    "net/http"
    "database/sql"
)

type Channel struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func createChannelHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var channel Channel
        decoder := json.NewDecoder(r.Body)
        if err := decoder.Decode(&channel); err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }
        defer r.Body.Close()

        // Insert the new channel into the database
        insertQuery := "INSERT INTO channels (name) VALUES ($1) RETURNING id, name"
        err := db.QueryRow(insertQuery, channel.Name).Scan(&channel.ID, &channel.Name)
        if err != nil {
            http.Error(w, "Error creating channel", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(channel)
    }
}
