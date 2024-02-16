package server

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/middleware"
	"github.com/gin-gonic/gin"
)

type HandlerOpts struct {
	Book         *handler.BookHandler
	User         *handler.UserHandler
	BorrowRecord *handler.BorrowRecordHandler
}

func SetupRouter(opts *HandlerOpts) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.HandleError)

	router.Use()
	router.PATCH("/borrows/:id", opts.BorrowRecord.ReturnBorrowedBook)
	router.POST("/login", opts.User.Login)

	router.Use(middleware.AuthHandler)

	router.GET("/books", opts.Book.GetBooks)
	router.GET("/users", opts.User.GetUsers)
	router.POST("/books", opts.Book.CreateBook)
	router.POST("/users", opts.User.CreateUser)
	router.POST("/borrows", opts.BorrowRecord.CreateBorrowRecord)
	return router
}
