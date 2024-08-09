package schema

import "time"

type (
	Blog struct {
		ID          uint64    `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
	}

	BlogWithUser struct {
		Blog
		CreatedBy User `json:"createdBy"`
	}

	BlogWithComment struct {
		Blog
		Comments []Comment `json:"comments"`
	}
)
