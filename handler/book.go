package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/usecase"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookUseCase usecase.BookUseCase
}

func NewBookHandler(bookUseCase usecase.BookUseCase) *BookHandler {
	return &BookHandler{
		bookUseCase: bookUseCase,
	}
}

func (h *BookHandler) GetAllBooks(ctx *gin.Context) {
	books, err := h.bookUseCase.GetAll()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Msg:  "OK",
		Data: books,
	})

}
