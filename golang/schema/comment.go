package schema

import "time"

type (
	Comment struct {
		ID          uint64    `json:"id"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
	}

	CommentWithBlog struct {
		Comment
		CreatedBy User `json:"createdBy,omitempty"`
		Blog      Blog `json:"blog"`
	}

	CommentWithReply struct {
		Comment
		CreatedBy User    `json:"createdBy,omitempty"`
		Blog      Blog    `json:"blog"`
		Replies   []Reply `json:"replies"`
	}
)
