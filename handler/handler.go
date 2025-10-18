package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hipostajm/crud-go/models"
	"github.com/hipostajm/crud-go/service"
	"github.com/hipostajm/crud-go/store"
)


type TaskHandler struct{
	TaskService service.TaskService
}



func (h* TaskHandler) Tasks(w http.ResponseWriter, r *http.Request){
	switch r.Method{
		case http.MethodGet:
			
			var input models.TaskReadInputModel
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil{
				http.Error(w, err.Error(), http.StatusBadRequest)
				fmt.Println(err)
				return
			}

			data, err := h.readTasks(input.UserID)
			if err != nil{
				http.Error(w, err.Error(), http.StatusBadRequest)
				fmt.Println(err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			io.Writer.Write(w, *data)

		case http.MethodPost:
			var input models.TaskCreateInputModel
			fmt.Println()
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil{
				http.Error(w, err.Error(), http.StatusBadRequest)
				fmt.Println(err)
				return
			}
			err := h.CreateTask(input)

			if err != nil{
				http.Error(w, err.Error(), http.StatusBadRequest)
				fmt.Println(err)
				return
			}

			io.WriteString(w, "sucess\n")

		case http.MethodPatch:
			var input models.TaskUpdateInputModel
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil{
				http.Error(w, err.Error(),  http.StatusBadRequest)
				fmt.Println(err)
				return
			}

			err := h.UpdateTask(input)

			if err != nil{
				http.Error(w, err.Error(),  http.StatusBadRequest)	
				fmt.Println(err)
				return
			}

			io.WriteString(w, "sucess\n")
		
		case http.MethodDelete:
			var input models.TaskDeleteInputModel
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil{
				http.Error(w, err.Error(),  http.StatusBadRequest)
				fmt.Println(err)
				return
			}

			err := h.DeleteTask(input)

			if err != nil{
				http.Error(w, err.Error(), http.StatusBadRequest)
				fmt.Println(err)
				return
			}

			io.WriteString(w, "sucess\n")

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h* TaskHandler) CreateTask(input models.TaskCreateInputModel) (error){
	task := &store.Task{CreatedAt: time.Now(), TaskName: input.TaskName, DueDate: input.DueDate, UserID: input.UserID}

	if input.Priority != nil{
		task.Priority = *input.Priority
	}

	if input.Status != nil{
		task.Status = *input.Status;
	}


	err := h.TaskService.CreateTask(task);

	return err
}

func (h* TaskHandler) readTasks(userID uint) (*[]byte, error){
	task, err := h.TaskService.ReadTasks(userID)
	if err != nil{
		return  nil, err
	}
	data, err :=  json.Marshal(task);
	if err != nil{
		return  nil, err
	}
	return &data, nil
}

func (h* TaskHandler) UpdateTask(input models.TaskUpdateInputModel) (error){
	task := &store.Task{UserID: input.UserID, ID: input.TaskID}

	if (input.DueDate != nil){
		task.DueDate = *input.DueDate
	}

	if (input.Priority != nil){
		task.Priority = *input.Priority
	}

	if (input.Status != nil){
		task.Status = *input.Status
	}

	if (input.TaskName != nil){
		task.TaskName = *input.TaskName
	}

	err := h.TaskService.UpdateTask(*task);
	return err
}

func (h* TaskHandler) DeleteTask(input models.TaskDeleteInputModel) (error){
	err := h.TaskService.DeleteTask(input.TaskID, input.UserID)
	return err
}