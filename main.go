package main

import (
	"log"
	"social-media-app/config"
	postHandler "social-media-app/feature/post/handler"
	postRepository "social-media-app/feature/post/repository"
	postService "social-media-app/feature/post/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	c := config.GetConfig()
	db := config.OpenDB(c)
	config.GormMigrartion(db)

	v := validator.New()

	// Setup feature
	// Post
	postRepo := postRepository.NewPostRepository(db)
	postSrv := postService.NewPostService(postRepo, v)
	postHandler := postHandler.NewPostHandler(postSrv)

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

	if err := e.Start(":8000"); err != nil {
		log.Fatal(err)
	}
}
