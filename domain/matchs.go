package domain

import "time"

// Matchs is representing the Article data struct
type Matchs struct {
	ID        int64     `json:"id"`
	UserOne   int64     `json:"user_one" validate:"required"`
	UserTwo   int64     `json:"user_two" validate:"required"`
	MatchTime time.Time `json:"match_time"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
