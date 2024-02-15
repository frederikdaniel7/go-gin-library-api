package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/constant"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/dto"
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			dto.Response{
				Msg:  err.Error(),
				Data: nil,
			})
		return
	}
	record, err := h.borrowRecordUseCase.NewBorrowRecord(body)
	if err != nil {
		if err.Error() == constant.ResponseMsgUserDoesNotExist {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				dto.Response{
					Msg:  err.Error(),
					Data: nil,
				})
			return
		}
		if err.Error() == constant.ResponseMsgBookDoesNotExist {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				dto.Response{
					Msg:  err.Error(),
					Data: nil,
				})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			dto.Response{
				Msg:  err.Error(),
				Data: nil,
			})
		return
	}
	ctx.JSON(http.StatusCreated, dto.Response{
		Msg:  "OK",
		Data: record,
	})

}
