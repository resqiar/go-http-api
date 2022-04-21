package main

import (
	"http-api/entities"
	"http-api/handler"
	"http-api/tasks"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// init db
	dsn := "host=localhost user=postgres password=admin dbname=db1 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	// auto migrate (must be off when on prod)
	db.AutoMigrate(&entities.User{}, &tasks.Task{})

	// init router
	r := gin.Default()

	v1 := r.Group("v1")
	// basic routes
	v1.GET("/hello", handler.HandleHelloRoute)
	v1.GET("/hello/:id", handler.HandleDetailRoute)
	v1.POST("/hello", handler.HandlePostHello)

	// TASK ROUTES
	taskRep := tasks.TaskRepository(db)
	taskService := tasks.TaskService(taskRep)
	taskCtrl := tasks.TaskController(taskService)

	v1.GET("/tasks", taskCtrl.HandleReadTask)
	v1.POST("/task/create", taskCtrl.HandleCreateTask)

	r.Run() // run on port 8080
}
