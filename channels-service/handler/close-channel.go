package handler

import (
    "net/http"
    "database/sql"
    "github.com/gorilla/mux"
)

func CloseChannelHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        params := mux.Vars(r)
        channelID := params["channelID"]

        // Implement logic to check if the channel exists
        checkQuery := "SELECT id FROM channels WHERE id = $1"
        var id int
        err := db.QueryRow(checkQuery, channelID).Scan(&id)
        if err != nil {
            if err == sql.ErrNoRows {
                http.Error(w, "Channel not found", http.StatusNotFound)
                return
            } else {
                http.Error(w, "Error checking the channel", http.StatusInternalServerError)
                return
            }
        }

        // Implement logic to close the channel with the given ID in the database
        deleteQuery := "DELETE FROM channels WHERE id = $1"
        _, err = db.Exec(deleteQuery, channelID)
        if err != nil {
            http.Error(w, "Error closing channel", http.StatusInternalServerError)
            return
        }

        // Return a response indicating successful channel closure
        w.WriteHeader(http.StatusNoContent)
    }
}
