package main

import (
	"http-api/auth"
	"http-api/guards"
	"http-api/tasks"
	"http-api/users"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Load env variables from .env file
	godotenv.Load()

	// Initialize Database (Gorm)
	// This should also be extracted to
	// env variables but for now it is OK.
	dsn := "host=localhost user=postgres password=admin dbname=db1 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	// Auto migration by Gorm.
	// Must be off when it comes to production.
	db.AutoMigrate(&users.User{}, &tasks.Task{})

	// Initialize gin engine
	// call it "r" as router.
	r := gin.Default()

	// Group all routes to have a prefix v1.
	// This is useful to maintan backward compatibility
	// when the API get updated in the future.
	v1 := r.Group("v1")

	// These are the module to initialize
	// controller and services for certain layers.
	// Although it feels weird and I doubt
	// this is the correct way to init all of 'em,
	// but idk the best practice in golang for now and it LGTM.
	userRep := users.UserRepository(db)
	userService := users.UserService(userRep)
	userCtrl := users.UserCtrl(userService)
	loginService := auth.LoginService(userService, userRep)
	loginCtrl := auth.LoginController(loginService, userService)
	taskRep := tasks.TaskRepository(db)
	taskService := tasks.TaskService(taskRep)
	taskCtrl := tasks.TaskController(taskService)

	// Group routes specifically for authentication
	// endpoint of the routes will be "v1/auth/..."
	authRoutes := v1.Group("auth")
	authRoutes.POST("/login", loginCtrl.Login)
	authRoutes.POST("/register", loginCtrl.Register)

	// User routes
	v1.GET("/users", userCtrl.HandleReadUsers)
	v1.GET("/user/:username", userCtrl.HandleFindUserByUsername)

	// Task routes
	v1.GET("/tasks", taskCtrl.HandleReadTask)
	v1.POST("/task/create", guards.JWTGuard(), taskCtrl.HandleCreateTask)

	r.Run() // run on port 8080 by default
}
