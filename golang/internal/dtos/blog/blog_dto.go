package blog_dtos

type (
	NewBlog struct {
		Title       string `json:"title" validate:"required"`
		Description string `json:"desc" validate:"required"`
	}

	UpdateBlog struct {
		ID          string `json:"id" validate:"required"`
		Title       string `json:"tile" validate:"required"`
		Description string `json:"desc" validate:"required"`
	}
)
