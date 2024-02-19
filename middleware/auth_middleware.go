package middleware

import (
	"net/http"
	"os"
	"strings"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/utils"
	"github.com/gin-gonic/gin"
)

type AuthOpts struct {
	Jwt utils.Crypto
}

func AuthHandler(c *gin.Context) {

	header := c.Request.Header.Get("Authorization")

	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Msg: "Unauthorized here",
		})
		return
	}

	token := strings.Split(header, " ")[1]
	claims, err := utils.ParseAndVerify(token, os.Getenv("SECRET_KEY"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Msg: err.Error(),
		})
		return
	}
	expired, err := claims.GetExpirationTime()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Msg: err.Error(),
		})
		return
	}
	if expired.Before(time.Now()) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Msg: err.Error(),
		})
		return
	}

	c.Set("id", claims["id"])
	c.Next()

}
