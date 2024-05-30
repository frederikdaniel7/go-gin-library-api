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

func (h *BookHandler) GetBooks(ctx *gin.Context) {

	title := ctx.Query("title")

	books, err := h.bookUseCase.GetBooks(ctx, title)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Msg: "OK",
		Data: dto.Books{
			Books: books,
		},
	})

}

func (h *BookHandler) CreateBook(ctx *gin.Context) {
	var body dto.CreateBookBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			dto.Response{
				Msg:  err.Error(),
				Data: nil,
			})
		return
	}
	book, err := h.bookUseCase.CreateBook(ctx, body)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, dto.Response{
		Msg:  "OK",
		Data: book,
	})
}
