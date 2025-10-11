package store

import (
	"context"
	"errors"
	"time"
	"gorm.io/gorm"
)

type Status string
const (
	PENDING Status = "Pending"
	INPROGRESS Status = "In-progress"
	COMPLETED Status = "Completed"
)

type Priority string
const (
	LOW Priority = "Low"
	MEDIUM Priority = "Medium"
	HIGH Priority = "High"
)

// To make those enums from Task struct
// CREATE TYPE status_enum AS ENUM ('Pending', 'In-Progress', 'Completed');
// CREATE TYPE priority_enum AS ENUM ('Low', 'Medium', 'High');


type Task struct{
	ID uint
	TaskName string
	UserID uint
	User User
	CreatedAt time.Time
	Status Status `gorm:"type:status_enum;default:'Pending'"`
	DueDate time.Time
	Priority Priority `gorm:"type:priority_enum;default:'Medium'"`
}

type User struct{
	ID uint 
	Username string `gorm:"uniqueIndex"`
	Role string
	Created_at time.Time
}

type TaskStore struct{
	Db *gorm.DB
	Ctx context.Context
}

type Store interface{
	CreateTask(name string, userId uint, date time.Time, status Status, dueDate time.Time, priority Priority) error
	ReadTask(taskId uint) (*Task, error)
	ReadTasksOfUser(userId uint) (*[]Task, error)
}

func (s *TaskStore) CreateTask(name string, userId uint, date time.Time, status Status, dueDate time.Time, priority Priority) error{
	err := s.Db.WithContext(s.Ctx).Create(&Task{TaskName: name, UserID: userId, CreatedAt: date, Status: status, DueDate: dueDate, Priority: priority}).Error
	return err
}

func (s *TaskStore) ReadTask(taskId uint) (*Task, error){
	return nil, errors.New("not implemented")
}

func (s *TaskStore) ReadTasksOfUser(userId uint) (*[]Task, error){
	return nil, errors.New("not implemented")
}