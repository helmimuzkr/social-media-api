package main

import (
	"log"
	"social-media-app/config"
	_postHandler "social-media-app/feature/post/handler"
	_postRepository "social-media-app/feature/post/repository"
	_postService "social-media-app/feature/post/service"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

func main() {
	c := config.GetConfig()
	db := config.OpenDB(c)
	config.GormMigrartion(db)

	v := validator.New()

	// Setup feature
	// Post
	postRepo := _postRepository.NewPostRepository(db)
	postSrv := _postService.NewPostService(postRepo, v)
	postHandler := _postHandler.NewPostHandler(postSrv)

	e := echo.New()

	if err := e.Start(":8000"); err != nil {
		log.Fatal(err)
	}
}
