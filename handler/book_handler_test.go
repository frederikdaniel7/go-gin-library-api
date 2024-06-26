package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/frederikdaniel7/go-gin-library-api/constant"
	"github.com/frederikdaniel7/go-gin-library-api/dto"
	"github.com/frederikdaniel7/go-gin-library-api/handler"
	"github.com/frederikdaniel7/go-gin-library-api/mocks"
	"github.com/frederikdaniel7/go-gin-library-api/server"
	"github.com/stretchr/testify/assert"
)

const (
	emptyTitle         = ""
	testTitle          = "test"
	testDescription    = "test description"
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

	t.Run("should return books when succesful", func(t *testing.T) {
		mockBookUseCase := new(mocks.BookUseCase)
		testCover := "test"
		testID := int64(1)
		testName := "test auth name"
		books := []dto.BookDetail{
			{
				ID:          1,
				Title:       "test",
				Description: "test",
				Quantity:    1,
				Cover:       &testCover,
				Author: &dto.Author{
					ID:   &testID,
					Name: &testName,
				},
			},
		}
		mockBookUseCase.On("GetBooks", emptyTitle).Return(books, nil)
		bookHandler := handler.NewBookHandler(mockBookUseCase)
		r := server.SetupRouter(&server.HandlerOpts{
			Book: bookHandler,
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/books", nil)
		r.ServeHTTP(w, req)
		response := dto.Response{
			Msg:  "OK",
			Data: books,
		}
		successMsg, _ := json.Marshal(response)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, string(successMsg), w.Body.String())
	})

	t.Run("should return books with certain title when succesful", func(t *testing.T) {
		mockBookUseCase := new(mocks.BookUseCase)
		testCover := "test"
		books := []dto.BookDetail{
			{
				ID:          1,
				Title:       "test",
				Description: "test",
				Quantity:    1,
				Cover:       &testCover,
				Author:      &dto.Author{},
			},
		}
		mockBookUseCase.On("GetBooks", testTitle).Return(books, nil)
		bookHandler := handler.NewBookHandler(mockBookUseCase)
		r := server.SetupRouter(&server.HandlerOpts{
			Book: bookHandler,
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/books", nil)
		q := req.URL.Query()
		q.Add("title", testTitle)
		req.URL.RawQuery = q.Encode()

		r.ServeHTTP(w, req)
		response := dto.Response{
			Msg:  "OK",
			Data: books,
		}
		successMsg, _ := json.Marshal(response)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, string(successMsg), w.Body.String())
	})

}

func TestBookHandler_CreateBook(t *testing.T) {
	t.Run("should be able to create book", func(t *testing.T) {
		testQuantity := 5
		body := dto.CreateBookBody{
			Title:       testTitle,
			Description: testDescription,
			Quantity:    &testQuantity,
		}
		book := dto.Book{
			ID:          1,
			Title:       testTitle,
			Description: testDescription,
			Quantity:    1,
			Cover:       nil,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
		}
		bodyJson, err := json.Marshal(body)
		if err != nil {
			log.Fatal(err.Error())
		}
		response := dto.Response{
			Msg:  constant.ResponseMsgOK,
			Data: book,
		}
		var expectedRes []byte
		expectedRes, err = json.Marshal(response)
		if err != nil {
			log.Fatal(err.Error())
		}
		mockUseCase := new(mocks.BookUseCase)
		mockUseCase.On("CreateBook", body).Return(&book, nil)
		bookHandler := handler.NewBookHandler(mockUseCase)

		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, "/books", bytes.NewReader(bodyJson))
		if err != nil {
			log.Fatal(err.Error())
		}
		router := server.SetupRouter(&server.HandlerOpts{
			Book: bookHandler,
		})
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Result().StatusCode)
		assert.Equal(t, string(expectedRes), w.Body.String())

	})
}
