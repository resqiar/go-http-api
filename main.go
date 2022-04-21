package main

import (
	"http-api/entities"
	"http-api/tasks"
	"http-api/users"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// init db
	dsn := "host=localhost user=postgres password=admin dbname=db1 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	// auto migrate (must be off when on prod)
	db.AutoMigrate(&entities.User{}, &tasks.Task{}, &users.User{})

	// init router
	r := gin.Default()

	v1 := r.Group("v1")

	// USER ROUTES
	userRep := users.UserRepository(db)
	userService := users.UserService(userRep)
	userCtrl := users.UserCtrl(userService)

	v1.GET("/users", userCtrl.HandleReadUsers)
	v1.POST("/user/create", userCtrl.HandleCreateUser)

	// TASK ROUTES
	taskRep := tasks.TaskRepository(db)
	taskService := tasks.TaskService(taskRep)
	taskCtrl := tasks.TaskController(taskService)

	v1.GET("/tasks", taskCtrl.HandleReadTask)
	v1.POST("/task/create", taskCtrl.HandleCreateTask)

	r.Run() // run on port 8080
}
