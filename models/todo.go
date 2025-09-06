package models

import ("time"
"gorm.io/gorm")

type TodoStatus string

const (
	StatusPending TodoStatus = "pending"
	StatusCompleted TodoStatus = "completed"
	StatusCancelled TodoStatus = "cancelled"
)

func (s TodoStatus) IsValid() bool {
	switch s {
	case StatusPending, StatusCompleted, StatusCancelled:
		return true
	default:
		return false
	}
}

type Todo struct {
	ID uint `json:"id" gorm:"primarykey"`
	Title string `json:"title" gorm:"not null;size:255" validate:"required,min=1,max=255"`
	Description string `json:"description" gorm:"size:2000"`
	Status TodoStatus `json:"status" gorm:"type:varchar(20);default:'pending'" validate:"required"`
	Priority int `json:"priority" gorm:"default:1" validate:"min=1, max=10"`
	DueDate *time.Time `json:"due_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type CreateTodoRequest struct {
	Title       string     `json:"title" validate:"required,min=1,max=255"`
	Description *string    `json:"description" validate:"omitempty,max=2000"`
	Priority    *int       `json:"priority" validate:"omitempty,min=1,max=10"`
	DueDate     *time.Time `json:"due_date"`
}


type UpdateTodoRequest struct {
	Title       *string     `json:"title" validate:"omitempty,min=1,max=255"`
	Description *string     `json:"description" validate:"omitempty,max=2000"`
	Status      *TodoStatus `json:"status" validate:"omitempty"`
	Priority    *int        `json:"priority" validate:"omitempty,min=1,max=10"`
	DueDate     *time.Time  `json:"due_date"`
}