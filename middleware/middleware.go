package middleware

import (
	"net/http"

	"github.com/frederikdaniel7/go-gin-library-api/constant"
	"github.com/frederikdaniel7/go-gin-library-api/dto"
	"github.com/frederikdaniel7/go-gin-library-api/exception"
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
