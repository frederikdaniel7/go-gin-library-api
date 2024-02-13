package main

import (
	"log"
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	InitDB()
	router := gin.Default()

	bookRepository := repository.NewBookRepository(db)

	bookUseCase := usecase.NewBookUseCaseImpl(bookRepository)

	bookHandler := handler.NewBookHandler(bookUseCase)

	router.GET("/test-db", func(ctx *gin.Context) {
		var num int
		err := db.QueryRow("SELECT 1").Scan(&num)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				dto.Response{
					Msg:  err.Error(),
					Data: nil,
				})
		}

		ctx.JSON(http.StatusOK,
			dto.Response{
				Msg:  "OK",
				Data: num,
			})
	})

	router.GET("/books", bookHandler.GetBooks)
	router.POST("/books", bookHandler.CreateBook)
	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}

}
