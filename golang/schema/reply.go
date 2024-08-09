package schema

import "time"

type (
	Reply struct {
		ID          uint64    `json:"id"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
		CreatedBy   User      `json:"createdBy"`
	}

	RepliesWthComment struct {
		Blog    Blog    `json:"blog"`
		Comment Comment `json:"comment"`
		Replies []Reply `json:"replies"`
	}
)
