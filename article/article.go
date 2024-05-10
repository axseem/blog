package article

import "time"

type Article struct {
	Title     string
	Content   string
	CreatedAt time.Time
}
