package store

import (
	"context"
	// "errors"
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
	CreateTask(task Task) (error)
	ReadTasks(userID uint) (*Task, error)
	ReadUser(userID uint) (User, error)
	ReadTask(taskID uint) (Task, error)
	UpdateTask(taskID uint, userID uint, task Task) (error)
    DeleteTask(taskID uint, userID uint) (error)
}

func (s *TaskStore) CreateTask(task *Task) error{
	err := s.Db.WithContext(s.Ctx).Create(task).Error
	return err
}

func (s *TaskStore) ReadTasks(userID uint) (*[]Task, error){
	var tasks []Task
	err := s.Db.WithContext(s.Ctx).Preload("User").Where("user_id = ?", userID).Find(&tasks).Error
	return &tasks, err
}

func (s *TaskStore) ReadUser(userID uint) (User, error){
	var user User
	err := s.Db.WithContext(s.Ctx).Where("id = ?", userID).First(&user).Error
	return user, err
}

func (s *TaskStore) ReadTask(taskID uint, userID uint) (Task, error){
	var task Task
	err := s.Db.WithContext(s.Ctx).Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error
	return task, err
}

func(s *TaskStore) UpdateTask(task Task) (error) {
	err := s.Db.WithContext(s.Ctx).Where("id = ? AND user_id = ?", task.ID, task.UserID).Updates(task).Error;
	return err
}

func(s *TaskStore) DeleteTask(taskID uint, userID uint) (error){
	err := s.Db.WithContext(s.Ctx).Where("id = ? AND user_id = ?", taskID, userID).Delete(&Task{}).Error
	return  err
}