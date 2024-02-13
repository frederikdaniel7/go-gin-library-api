package server

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/handler"
	"github.com/gin-gonic/gin"
)

type HandlerOpts struct {
	Book *handler.BookHandler
}

func SetupRouter(opts *HandlerOpts) *gin.Engine {
	router := gin.Default()

	router.GET("/books", opts.Book.GetBooks)
	router.POST("/books", opts.Book.CreateBook)
	return router
}
