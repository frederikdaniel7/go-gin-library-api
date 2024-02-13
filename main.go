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

	bookUseCase := usecase.NewBookUseCaseImpl(bookRepository)

	bookHandler := handler.NewBookHandler(bookUseCase)

	router := server.SetupRouter(&server.HandlerOpts{
		Book: bookHandler,
	})

	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}

}
