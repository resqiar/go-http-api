package main

import (
	"fmt"
	"http-api/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=admin dbname=exampledb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	_, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		log.Fatal(dbErr)
	} else {
		fmt.Println("...Postgres connected...")
	}

	r := gin.Default()

	v1 := r.Group("v1")
	// basic routes
	v1.GET("/hello", handler.HandleHelloRoute)
	v1.GET("/hello/:id", handler.HandleDetailRoute)
	v1.POST("/hello", handler.HandlePostHello)

	r.Run() // run on port 8080
}
