package model

import "time"

type Article struct {
	ID        string
	Title     string
	Content   string
	CreatedAt time.Time
}
