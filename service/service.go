package service

import(
	"github.com/hipostajm/crud-go/store"
	"time"
)

type TaskService struct{
	TaskStore store.TaskStore	
}

func (s* TaskService) CreateTask() error{
	err :=s.TaskStore.CreateTask("qt", 1, time.Now(), store.PENDING, time.Now(), store.MEDIUM)
	return err
}

func (s* TaskService) ReadTask(){

}

func (s* TaskService) UpdateTask(){

}

func (s* TaskService) DeleteTask(){

}