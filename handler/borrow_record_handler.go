package handler

import (
	"fmt"
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/exception"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/usecase"
	"github.com/gin-gonic/gin"
)

type BorrowRecordHandler struct {
	borrowRecordUseCase usecase.BorrowRecordUseCase
}

func NewBorrowRecordHandler(borrowRecordUseCase usecase.BorrowRecordUseCase) *BorrowRecordHandler {
	return &BorrowRecordHandler{
		borrowRecordUseCase: borrowRecordUseCase,
	}
}

func (h *BorrowRecordHandler) CreateBorrowRecord(ctx *gin.Context) {

	var body dto.CreateBorrowRecordBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(err)
		return
	}
	record, err := h.borrowRecordUseCase.NewBorrowRecord(ctx, body)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, dto.Response{
		Msg:  "OK",
		Data: record,
	})

}

func (h *BorrowRecordHandler) ReturnBorrowedBook(ctx *gin.Context) {
	var idParam dto.ReturnBookParam
	err := ctx.ShouldBindUri(&idParam)
	if err != nil {
		exception.NewErrorType(http.StatusBadRequest, fmt.Sprintf("%v", idParam))
		return
	}
	record, err := h.borrowRecordUseCase.ReturnBorrowedBook(ctx, idParam.ID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusAccepted, dto.Response{
		Msg:  "OK",
		Data: record,
	})

}
