// message.go

package models

type Message struct {
    ID        int    `json:"id"`
    Content   string `json:"content"`
    // Add more fields for other message types (image, audio, video, file, etc.)
}
