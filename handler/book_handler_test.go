package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/server"
	"github.com/stretchr/testify/assert"
)

const (
	emptyTitle         = ""
	testTitle          = "test"
	errorReturnNothing = "error return nothing"
)

func TestBookHandler_GetBooks(t *testing.T) {
	t.Run("should return error no books", func(t *testing.T) {
		mockBookUseCase := new(mocks.BookUseCase)

		mockBookUseCase.On("GetBooks", emptyTitle).Return(nil, errors.New(errorReturnNothing))
		bookHandler := handler.NewBookHandler(mockBookUseCase)
		r := server.SetupRouter(&server.HandlerOpts{
			Book: bookHandler,
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/books", nil)

		r.ServeHTTP(w, req)
		response := dto.Response{Msg: errorReturnNothing}

		errMsg, _ := json.Marshal(response)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
		assert.Equal(t, string(errMsg), w.Body.String())
	})

}
