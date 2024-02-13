package main

import (
	"log"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/server"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/usecase"
)

func main() {
	InitDB()

	bookRepository := repository.NewBookRepository(db)
	userRepository := repository.NewUserRepository(db)
	authorRepository := repository.NewAuthorRepository(db)

	bookUseCase := usecase.NewBookUseCaseImpl(bookRepository, authorRepository)
	userUseCase := usecase.NewUserUseCaseImpl(userRepository)

	bookHandler := handler.NewBookHandler(bookUseCase)
	userHandler := handler.NewUserHandler(userUseCase)
	router := server.SetupRouter(&server.HandlerOpts{
		Book: bookHandler,
		User: userHandler,
	})

	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}

}
