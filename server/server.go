package server

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/handler"
	"github.com/gin-gonic/gin"
)

type HandlerOpts struct {
	Book *handler.BookHandler
	User *handler.UserHandler
}

func SetupRouter(opts *HandlerOpts) *gin.Engine {
	router := gin.Default()

	router.GET("/books", opts.Book.GetBooks)
	router.GET("/users", opts.User.GetUsers)
	router.POST("/books", opts.Book.CreateBook)
	return router
}
