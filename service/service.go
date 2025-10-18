package service

import (
	"errors"

	"github.com/hipostajm/crud-go/store"
)

type TaskService struct{
	TaskStore store.TaskStore	
}


// ID uint
// TaskName string
// UserID uint
// User User
// CreatedAt time.Time
// Status Status `gorm:"type:status_enum;default:'Pending'"`
// DueDate time.Time
// Priority Priority `gorm:"type:priority_enum;default:'Medium'"`

func (s* TaskService) ValidateUser(userID uint) (bool){
	_, err := s.TaskStore.ReadUser(userID)
	return err != nil
}

func (s* TaskService) ValidateTask(taskID uint, userID uint) (bool){
	_, err := s.TaskStore.ReadTask(taskID, userID)
	return err != nil
}

func (s* TaskService) CreateTask(task *store.Task) error{
	
	if task.CreatedAt.Compare(task.DueDate) == 1 {
		return  errors.New("create date is smaller or equel then dueDate")
	}

	if task.TaskName == ""{
		return errors.New("no task name")
	}


	if s.ValidateUser(task.UserID){
		return errors.New("bad user id")
	}

	// if (task.Priority != store.HIGH) || (task.Priority != store.MEDIUM) || (task.Priority != store.LOW){
	// 	return errors.New("bad priority")
	// }

	// if (task.Status != store.PENDING) || (task.Status != store.COMPLETED) || (task.Status != store.INPROGRESS){
	// 	return errors.New("bad status")
	// }

	err := s.TaskStore.CreateTask(task);
	return err
}

func (s* TaskService) ReadTasks(userID uint) (*[]store.Task, error){

	if s.ValidateUser(userID){
		return nil, errors.New("bad user id")
	}


	task, err := s.TaskStore.ReadTasks(userID);

	return task, err;

}

func (s* TaskService) UpdateTask(task store.Task) (error){
	if s.ValidateUser(task.UserID){
		return errors.New("bad user id")
	}

	if s.ValidateTask(task.ID, task.UserID){
		return errors.New("bad task id or user does not created the task")
	}

	err := s.TaskStore.UpdateTask(task)

	return err
}

func (s* TaskService) DeleteTask(taskID uint, userID uint)(error){

	if s.ValidateUser(userID){
		return errors.New("bad user id")
	}

	if s.ValidateTask(taskID, userID){
		return errors.New("bad task id or user does not created the task")
	}

	err := s.TaskStore.DeleteTask(taskID, userID)
	return err
}