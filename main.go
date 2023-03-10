package main

import (
	"log"
	"social-media-app/config"
	postHandler "social-media-app/feature/post/handler"
	postRepository "social-media-app/feature/post/repository"
	postService "social-media-app/feature/post/service"

	_commentHandler "social-media-app/feature/comment/handler"
	_commentRepository "social-media-app/feature/comment/repository"
	_commentService "social-media-app/feature/comment/service"

	userHandler "social-media-app/feature/user/handler"
	userRepository "social-media-app/feature/user/repository"
	userService "social-media-app/feature/user/service"

	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	c := config.GetConfig()
	db := config.OpenDB(c)
	config.GormMigration(db)

	v := validator.New()

	// Setup feature
	postRepo := postRepository.NewPostRepository(db)
	postSrv := postService.NewPostService(postRepo, v)
	postHandler := postHandler.NewPostHandler(postSrv)
	commentRepo := _commentRepository.NewCommentRepository(db)
	commentSrv := _commentService.NewCommentService(commentRepo, v)
	commentHandler := _commentHandler.NewCommentHandler(commentSrv)
	userRepo := userRepository.New(db)
	userSrv := userService.New(userRepo, v)
	userHandler := userHandler.New(&userSrv)

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom}, method=${method}, uri=${uri}, status=${status}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))

	e.POST("/posts", postHandler.Create(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/posts", postHandler.MyPost(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/posts/:post_id", postHandler.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/posts/:post_id", postHandler.Delete(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/posts/:post_id", postHandler.GetByID())
	e.GET("/posts/list/:user_id", postHandler.GetByUserID())
	e.GET("/posts/list", postHandler.GetAll())

	e.POST("/comments", commentHandler.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/comments", commentHandler.GetAll())

	e.POST("/register", userHandler.RegisterHand())
	e.POST("/login", userHandler.LoginHand())

	userLogin := e.Group("/users")
	userLogin.Use(middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/search", userHandler.SearchHand(), middleware.JWT([]byte(config.JWT_KEY)))
	userLogin.GET("/:id", userHandler.GetByIdHand())
	userLogin.GET("", userHandler.ProfileHand())
	userLogin.PUT("", userHandler.UpdateHand())
	userLogin.PUT("/password", userHandler.UpdatePassHand())
	userLogin.DELETE("", userHandler.RemoveHand())

	if err := e.Start(":8000"); err != nil {
		log.Fatal(err)
	}
}
