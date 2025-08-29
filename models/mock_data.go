package models

import "time"

var (
	Todo1 = Todo{
		ID:        1,
		Title:     "todo1 title",
		Content:   "todo1 content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	Todo2 = Todo{
		ID:        2,
		Title:     "todo2 title",
		Content:   "todo2 content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
)
