package usecase

import (
	"context"
	"net/http"

	"github.com/frederikdaniel7/go-gin-library-api/constant"
	"github.com/frederikdaniel7/go-gin-library-api/dto"
	"github.com/frederikdaniel7/go-gin-library-api/entity"
	"github.com/frederikdaniel7/go-gin-library-api/exception"
	"github.com/frederikdaniel7/go-gin-library-api/repository"
	"github.com/frederikdaniel7/go-gin-library-api/utils"
)

type BookUseCase interface {
	GetBooks(ctx context.Context, query dto.BookQuery) (*dto.BookResponse, error)
	CreateBook(ctx context.Context, body dto.CreateBookBody) (*dto.Book, error)
}

type bookUseCaseImpl struct {
	bookRepository   repository.BookRepository
	authorRepository repository.AuthorRepository
}

func NewBookUseCaseImpl(bookRepository repository.BookRepository, authorRepository repository.AuthorRepository) *bookUseCaseImpl {
	return &bookUseCaseImpl{
		bookRepository:   bookRepository,
		authorRepository: authorRepository,
	}
}

func (b *bookUseCaseImpl) GetBooks(ctx context.Context, query dto.BookQuery) (*dto.BookResponse, error) {
	var result dto.BookResponse
	booksJson := []dto.BookDetail{}

	var books []entity.BookDetail
	var bookPackage entity.BookPackage
	var err error
	if query.Size == 0 {
		query.Size = constant.DefaultPaginationSize
	}
	if query.Page == 0 {
		query.Page = constant.DefaultPaginationPage
	}

	if query.Title == "" {
		res, err := b.bookRepository.FindAll(ctx, entity.Paging{
			Size: &query.Size,
			Page: &query.Page,
		})

		if err == nil {
			bookPackage = *res
		}
	} else {
		books, err = b.bookRepository.FindSimilarBookByTitle(ctx, query.Title)
		if err != nil && len(books) != 0 {
			bookPackage.Books = books
		}
	}
	if err != nil {
		return nil, exception.NewErrorType(http.StatusBadRequest, err.Error())
	}

	for _, book := range bookPackage.Books {
		booksJson = append(booksJson, utils.ConvertBookDetailToJson(book))
	}

	result.Books = booksJson
	result.ItemCount = bookPackage.Count
	result.PageCount = (bookPackage.TotalData + query.Size - 1) / query.Size
	result.CurrentPage = query.Page

	return &result, nil

}

func (b *bookUseCaseImpl) CreateBook(ctx context.Context, body dto.CreateBookBody) (*dto.Book, error) {
	if body.AuthorID != nil {
		checkAuthorExists, err := b.authorRepository.FindOneById(ctx, *body.AuthorID)
		if checkAuthorExists.ID == nil {
			return nil, exception.NewErrorType(
				http.StatusNotFound,
				constant.ResponseMsgAuthorDoesNotExist)
		}
		if err != nil {
			return nil, err
		}
	}

	checkExist, err := b.bookRepository.FindSimilarBookByTitle(ctx, body.Title)
	if err != nil {
		return nil, err
	}
	for _, book := range checkExist {
		if body.Title == book.Title {
			return nil, exception.NewErrorType(
				http.StatusPreconditionFailed,
				constant.ResponseMsgBookAlreadyExists)
		}
	}

	book, err := b.bookRepository.CreateBook(ctx, body)
	if err != nil {
		return nil, err
	}
	bookJson := utils.ConvertBookToJson(*book)
	return &bookJson, nil
}
