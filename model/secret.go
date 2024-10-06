package model

import "time"

type Secret struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Note      string     `json:"note"`
	Email     string     `json:"email"`
	Website   string     `json:"website"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
