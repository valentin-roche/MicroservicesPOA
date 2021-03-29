package blog

import (
	"time"
)

type BlogPost struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Content     string    `json:"content"`
	PublishedOn time.Time `json:"published_on"`
}
