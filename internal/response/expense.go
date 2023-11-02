package response

import "time"

type CommonExpenseResponse struct {
	ID          int       `json:"id"`
	UserID      uint      `json:"userId"`
	AccountID   uint      `json:"accountId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
