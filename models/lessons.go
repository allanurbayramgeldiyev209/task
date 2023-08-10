package models

type Lesson struct {
	ID          string `json:"id,omitempty"`
	OrderNumber int    `json:"order_number,omitempty"`
	Name        string `json:"name,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	DayNumber   int    `json:"day_number,omitempty"`
	CreatedAt   string `json:"-"`
	UpdatedAt   string `json:"-"`
	DeletedAt   string `json:"-"`
}
