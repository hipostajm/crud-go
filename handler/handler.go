package handler

import (
	"net/http"
	"io"
	"github.com/hipostajm/crud-go/service"
)

type TaskHandler struct{
	TaskService service.TaskService
}

func (h* TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request){
	err := h.TaskService.CreateTask()
	if err != nil{
		panic(err)
	}
	io.WriteString(w, "this works?\n")
}

func (h* TaskHandler) ReadTask(w http.ResponseWriter, r *http.Request){

}

func (h* TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request){

}

func (h* TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request){
}