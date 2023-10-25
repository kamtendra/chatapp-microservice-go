// router/routes.go

package router

import (
	"chat-app-microservice/channels-service/model"
	"chat-app-microservice/channels-service/repository"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func SetupRoutes(r *mux.Router, db *sql.DB) *mux.Router {
	channelRepo := repository.NewChannelRepository(db)

	r.HandleFunc("/channels", createChannelHandler(channelRepo)).Methods("POST")
	r.HandleFunc("/channels/{channelID:[0-9]+}", closeChannelHandler(channelRepo)).Methods("DELETE")
	r.HandleFunc("/channels", queryChannelsHandler(channelRepo)).Methods("GET")

	return r
}

func createChannelHandler(repo *repository.ChannelRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var channel model.Channel
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&channel); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Insert the new channel into the database
		channelID, err := repo.CreateChannel(channel.Name)
		if err != nil {
			http.Error(w, "Error creating channel", http.StatusInternalServerError)
			return
		}

		// Return a JSON response with the newly created channel
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(struct {
			ChannelID int `json:"channel_id"`
		}{ChannelID: channelID})
	}
}

func closeChannelHandler(repo *repository.ChannelRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		channelID, err := strconv.Atoi(params["channelID"])
		if err != nil {
			http.Error(w, "Invalid channel ID", http.StatusBadRequest)
			return
		}

		// Implement the logic to close the channel with the given channelID in the database
		err = repo.CloseChannel(channelID)
		if err != nil {
			http.Error(w, "Error closing channel", http.StatusInternalServerError)
			return
		}

		// Return a response indicating successful channel closure
		w.WriteHeader(http.StatusNoContent)
	}
}

func queryChannelsHandler(repo *repository.ChannelRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve a list of channels from the database
		channels, err := repo.GetChannels()
		if err != nil {
			http.Error(w, "Error querying channels", http.StatusInternalServerError)
			return
		}

		// Return the list of channels as a JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(channels)
	}
}
