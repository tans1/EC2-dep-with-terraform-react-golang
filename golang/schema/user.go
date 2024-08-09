package schema

import "time"

type (
	User struct {
		ID          uint64    `json:"id"`
		FirstName   string    `json:"firstName"`
		LastName    string    `json:"lastName"`
		Username    string    `json:"username"`
		Email       string    `json:"email,omitempty"`
		PhoneNumber string    `json:"phoneNumber"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
	}

	UserWithBlog struct {
		User
		Blogs []Blog `json:"blogs"`
	}

	UserWithPassword struct {
		User
		Password string `json:"password"`
	}
)
