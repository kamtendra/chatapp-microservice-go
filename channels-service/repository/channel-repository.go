// repository/channel_repository.go

package repository

import (
    "database/sql"
    "log"
    "fmt"
)

type ChannelRepository struct {
    DB *sql.DB
}

func NewChannelRepository(db *sql.DB) *ChannelRepository {
    return &ChannelRepository{DB: db}
}

func (repo *ChannelRepository) CreateChannel(name string) (int, error) {
    var channelID int
    err := repo.DB.QueryRow("INSERT INTO channels (name) VALUES ($1) RETURNING id", name).Scan(&channelID)
    if err != nil {
        log.Println("Error creating channel:", err)
        return 0, err
    }
    return channelID, nil
}

func (repo *ChannelRepository) CloseChannel(channelID int) error {
    _, err := repo.DB.Exec("DELETE FROM channels WHERE id = $1", channelID)
    if err != nil {
        log.Println("Error closing channel:", err)
        return err
    }
    return nil
}

func (repo *ChannelRepository) GetChannels() ([]model.Channel, error) {
    rows, err := repo.DB.Query("SELECT id, name FROM channels")
    if err != nil {
        log.Println("Error querying channels:", err)
        return nil, err
    }
    defer rows.Close()

    var channels []model.Channel
    for rows.Next() {
        var channel model.Channel
        err := rows.Scan(&channel.ID, &channel.Name)
        if err != nil {
            log.Println("Error scanning channel:", err)
            return nil, err
        }
        channels = append(channels, channel)
    }
    return channels, nil
}
