package model

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	Nome      string    `json:"nome" db:"nome"`
	Idade     int       `json:"idade" db:"idade"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
