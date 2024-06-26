package handler

import (
	"fmt"
	"net/http"

	"github.com/frederikdaniel7/go-gin-library-api/dto"
	"github.com/frederikdaniel7/go-gin-library-api/exception"
	"github.com/frederikdaniel7/go-gin-library-api/usecase"
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
	userId := ctx.GetFloat64("id")
	var body dto.CreateBorrowRecordBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Error(err)
		return
	}
	record, err := h.borrowRecordUseCase.NewBorrowRecord(ctx, body, int(userId))
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

	ctx.JSON(http.StatusOK, dto.Response{
		Msg:  "OK",
		Data: record,
	})

}
