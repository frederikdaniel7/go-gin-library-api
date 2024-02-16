package usecase

import (
	"context"
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/constant"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/exception"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/utils"
)

type BookUseCase interface {
	GetBooks(ctx context.Context, title string) ([]dto.BookDetail, error)
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

func (b *bookUseCaseImpl) GetBooks(ctx context.Context, title string) ([]dto.BookDetail, error) {

	booksJson := []dto.BookDetail{}
	var books []entity.BookDetail
	var err error
	if title == "" {

		books, err = b.bookRepository.FindAll(ctx)

	} else {
		books, err = b.bookRepository.FindSimilarBookByTitle(ctx, title)

	}
	if err != nil {
		return nil, exception.NewErrorType(http.StatusBadRequest, err.Error())
	}
	for _, book := range books {
		booksJson = append(booksJson, utils.ConvertBookDetailToJson(book))
	}

	return booksJson, nil

}

func (b *bookUseCaseImpl) CreateBook(ctx context.Context, body dto.CreateBookBody) (*dto.Book, error) {

	checkAuthorExists, err := b.authorRepository.FindOneById(ctx, *body.AuthorID)
	if checkAuthorExists.ID == nil {
		return nil, exception.NewErrorType(
			http.StatusNotFound,
			constant.ResponseMsgAuthorDoesNotExist)
	}
	if err != nil {
		return nil, err
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
