package server

import (
	"net/http"
	"net/http/pprof"

	"github.com/frederikdaniel7/go-gin-library-api/handler"
	"github.com/frederikdaniel7/go-gin-library-api/middleware"
	"github.com/gin-gonic/gin"
)

type HandlerOpts struct {
	Book         *handler.BookHandler
	User         *handler.UserHandler
	BorrowRecord *handler.BorrowRecordHandler
}

func SetupRouter(opts *HandlerOpts) *gin.Engine {
	router := gin.Default()
	router.ContextWithFallback = true
	router.Use(middleware.HandleError)

	router.Use()

	router.POST("/users", opts.User.CreateUser)
	router.POST("/login", opts.User.Login)
	router.GET("/books", opts.Book.GetBooks)
	router.GET("/debug/pprof/", gin.WrapH(http.HandlerFunc(pprof.Index)))

	router.GET("/debug/pprof/profile", gin.WrapH(http.HandlerFunc(pprof.Profile)))

	router.GET("/debug/pprof/heap", gin.WrapH(http.HandlerFunc(pprof.Handler("heap").ServeHTTP)))

	router.GET("/debug/pprof/block", gin.WrapH(http.HandlerFunc(pprof.Handler("block").ServeHTTP)))

	router.GET("/debug/pprof/goroutine", gin.WrapH(http.HandlerFunc(pprof.Handler("goroutine").ServeHTTP)))

	router.Use(middleware.AuthHandler)
	router.PATCH("/borrows/:id", opts.BorrowRecord.ReturnBorrowedBook)

	router.GET("/users", opts.User.GetUsers)
	router.POST("/books", opts.Book.CreateBook)
	router.POST("/borrows", opts.BorrowRecord.CreateBorrowRecord)
	return router
}
