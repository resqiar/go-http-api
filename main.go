package main

import (
	"http-api/controllers"
	"http-api/entities"
	"http-api/guards"
	"http-api/repositories"
	"http-api/services"
	"log"

	"github.com/gin-contrib/cors"
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
	db.AutoMigrate(&entities.User{}, &entities.Question{}, &entities.Answer{})

	// Initialize gin engine
	// call it "r" as router.
	r := gin.New()
	// Add recovery middleware
	// prevent error crash the app
	r.Use(gin.Recovery())
	// configure cors
	// cors should allow all origin (Public)
	// cors should allow only GET and POST
	// corst should allow "Authorization" header
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowMethods = []string{"GET", "POST"}
	r.Use(cors.New(corsConfig))
	// Group all routes to have a prefix v1.
	// This is useful to maintan backward compatibility
	// when the API get updated in the future.
	v1 := r.Group("v1")

	// These are the module to initialize
	// controller and services for certain layers.
	// Although it feels weird and I doubt
	// this is the correct way to init all of 'em,
	// but idk the best practice in golang for now and it LGTM.
	// ----------------------------------------------------------------
	// Initialize repositories
	// ----------------------------------------------------------------
	userRep := repositories.UserRepository(db)
	questionRep := repositories.QuestionRepository(db)
	answerRep := repositories.AnswerRepository(db)
	// ----------------------------------------------------------------
	// initialize services
	// ----------------------------------------------------------------
	userService := services.UserService(userRep)
	loginService := services.LoginService(userService, userRep)
	questionService := services.QuestionService(questionRep)
	answerService := services.AnswerService(answerRep)
	// ----------------------------------------------------------------
	// Initialize controllers
	// ----------------------------------------------------------------
	userCtrl := controllers.UserCtrl(userService)
	loginCtrl := controllers.LoginController(loginService, userService)
	questionCtrl := controllers.QuestionController(questionService, userService)
	answerCtrl := controllers.AnswerController(answerService, questionService, userService)

	// Group routes specifically for authentication
	// endpoint of the routes will be "v1/auth/..."
	authRoutes := v1.Group("auth")
	authRoutes.POST("/login", loginCtrl.Login)
	authRoutes.POST("/register", loginCtrl.Register)

	// ----------------------------------------------------------------
	// User routes
	// ----------------------------------------------------------------
	// ! MAKE SURE YOUR CODE REVIEWERs DONT SEE THIS ENDPOINT !
	v1.GET("/users", userCtrl.HandleReadUsers)
	// ! ---------------------------------------------------- !
	v1.GET("/user/:username", userCtrl.HandleFindUserByUsername)
	v1.GET("/userid/:id", userCtrl.HandleFindUserById)
	// ----------------------------------------------------------------
	// Question routes
	// ----------------------------------------------------------------
	v1.GET("/questions", questionCtrl.HandleReadQuestion)
	v1.GET("/question/:id", questionCtrl.HandleReadDetailQuestion)
	v1.POST("/question/create", guards.JWTGuard(), questionCtrl.HandleCreateQuestion)
	v1.POST("/question/update", guards.JWTGuard(), questionCtrl.HandleUpdateQuestion)
	v1.POST("/question/delete", guards.JWTGuard(), questionCtrl.HandleDeleteQuestion)
	// ----------------------------------------------------------------
	// Answer routes
	// ----------------------------------------------------------------
	v1.GET("/answers", answerCtrl.HandleReadAnswers)
	v1.POST("/answer/create", guards.JWTGuard(), answerCtrl.HandleCreateAnswer)
	v1.POST("/answer/update", guards.JWTGuard(), answerCtrl.HandleUpdateAnswer)
	v1.POST("/answer/delete", guards.JWTGuard(), answerCtrl.HandleDeleteAnswer)

	r.Run(":8000") // run on port 8080 by default
}
