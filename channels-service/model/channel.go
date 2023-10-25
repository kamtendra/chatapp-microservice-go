
package model

type Channel struct {
    ID          int    `json:"id"`
    Participants []int  `json:"participants"` // User IDs of participants
    CreatedAt   string `json:"created_at"`   // Creation date, you can use a string or a time.Time
    IsClosed    bool   `json:"is_closed"`    // Closure status
}
