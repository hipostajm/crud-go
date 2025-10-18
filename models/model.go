package models

import(
	"github.com/hipostajm/crud-go/store"
	"time"
)

// ID uint
// TaskName string
// UserID uint
// User User
// CreatedAt time.Time
// Status Status `gorm:"type:status_enum;default:'Pending'"`
// DueDate time.Time
// Priority Priority `gorm:"type:priority_enum;default:'Medium'"`

type TaskCreateInputModel struct {
	TaskName string `json:"task_name"`
	UserID uint `json:"user_id"`
	DueDate time.Time `json:"due_date"`
	Status *store.Status `json:"status"`
	Priority *store.Priority `json:"priority"`
}

type TaskReadInputModel struct{
	UserID uint `json:"user_id"`
}

type TaskUpdateInputModel struct{
	TaskID uint `json:"task_id"`
	UserID uint `json:"user_id"`
	Status *store.Status `json:"status"`
	Priority *store.Priority `json:"priority"`
	TaskName *string `json:"task_name"`
	DueDate *time.Time `json:"due_date"`
}

type TaskDeleteInputModel struct{
	UserID uint `json:"user_id"`
	TaskID uint `json:"task_id"`
}