package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) GetUsers(ctx *gin.Context) {
	name := ctx.Query("name")

	users, err := h.userUseCase.GetUsers(name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			dto.Response{
				Msg:  err.Error(),
				Data: nil,
			})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Msg:  "OK",
		Data: users,
	})
}
