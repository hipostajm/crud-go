package main

import (
	"context"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/hipostajm/crud-go/handler"
	"github.com/hipostajm/crud-go/service"
	"github.com/hipostajm/crud-go/store"
)

func main() {
	dsn := "host=localhost user=admin password=admin dbname=db port=5431 sslmode=disable TimeZone=Europe/Warsaw"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil{
    panic("db made a error")
  }

  ctx := context.Background()

  db.AutoMigrate(&store.Task{}, &store.User{})

  taskStore := store.TaskStore{Ctx: ctx, Db: db}
  taskServiec := service.TaskService{TaskStore: taskStore}
  taskHandler := handler.TaskHandler{TaskService: taskServiec}

  http.HandleFunc("/create", taskHandler.CreateTask)

  if err := http.ListenAndServe(":8080", nil); err != nil{
    panic(err)
  }
}