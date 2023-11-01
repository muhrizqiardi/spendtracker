package response

import "time"

type CommonCategoryResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"userId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
