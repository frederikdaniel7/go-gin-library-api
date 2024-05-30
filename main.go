package main

import (
	"log"

	"github.com/frederikdaniel7/go-gin-library-api/database"
	"github.com/frederikdaniel7/go-gin-library-api/handler"
	"github.com/frederikdaniel7/go-gin-library-api/repository"
	"github.com/frederikdaniel7/go-gin-library-api/server"
	"github.com/frederikdaniel7/go-gin-library-api/usecase"
)

func main() {
	if err := ConfigInit(); err != nil {
		log.Fatalf("error env : %s", err.Error())
	}
	InitDB()

	bookRepository := repository.NewBookRepository(db)
	userRepository := repository.NewUserRepository(db)
	authorRepository := repository.NewAuthorRepository(db)
	borrowRecordRepository := repository.NewBorrowRecordRepository(db)
	transactor := database.NewTransaction(db)

	bookUseCase := usecase.NewBookUseCaseImpl(bookRepository, authorRepository)
	userUseCase := usecase.NewUserUseCaseImpl(userRepository)
	borrowRecordUseCase := usecase.NewBorrowRecordUseCaseImpl(
		borrowRecordRepository, bookRepository, userRepository, transactor)

	bookHandler := handler.NewBookHandler(bookUseCase)
	userHandler := handler.NewUserHandler(userUseCase)
	borrowRecordHandler := handler.NewBorrowRecordHandler(borrowRecordUseCase)
	router := server.SetupRouter(&server.HandlerOpts{
		Book:         bookHandler,
		User:         userHandler,
		BorrowRecord: borrowRecordHandler,
	})

	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}

}
