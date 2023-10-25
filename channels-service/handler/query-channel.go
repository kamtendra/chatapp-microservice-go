package handler

import (
    "encoding/json"
    "net/http"
    "database/sql"
    "chat-app-microservice/channels-service/model"
)

func QueryChannelsHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        rows, err := db.Query("SELECT id, created_at, is_closed FROM channels")
        if err != nil {
            http.Error(w, "Error querying channels", http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        var channels []model.Channel
        for rows.Next() {
            var channel model.Channel
            if err := rows.Scan(&channel.ID, &channel.CreatedAt, &channel.IsClosed); err != nil {
                http.Error(w, "Error scanning channel", http.StatusInternalServerError)
                return
            }
            channels = append(channels, channel)
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(channels); err != nil {
            http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
        }
    }
}
