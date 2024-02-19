package middleware

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/constant"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/exception"
	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context) {

	c.Next()

	for _, err := range c.Errors {
		switch e := err.Err.(type) {
		case *exception.ErrorType:
			{
				c.AbortWithStatusJSON(e.StatusCode,
					dto.Response{
						Msg:  e.Message,
						Data: nil,
					})
			}
		default:
			{
				c.AbortWithStatusJSON(http.StatusInternalServerError,
					dto.Response{
						Msg:  constant.ResponseMsgErrorInternal,
						Data: nil,
					})
			}

		}
	}
	
}
