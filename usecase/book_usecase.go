package usecase

import (
	"errors"
	"fmt"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/constant"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/utils"
)

type BookUseCase interface {
	GetBooks(title string) ([]dto.BookDetail, error)
	CreateBook(body dto.CreateBookBody) (*dto.Book, error)
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

func (b *bookUseCaseImpl) GetBooks(title string) ([]dto.BookDetail, error) {

	booksJson := []dto.BookDetail{}
	var books []entity.BookDetail
	var err error
	if title == "" {
		books, err = b.bookRepository.FindAll()

	} else {
		books, err = b.bookRepository.FindSimilarBookByTitle(title)

	}
	if err != nil {
		return nil, err
	}
	for _, book := range books {
		booksJson = append(booksJson, utils.ConvertBookDetailToJson(book))
	}

	return booksJson, nil

}

func (b *bookUseCaseImpl) CreateBook(body dto.CreateBookBody) (*dto.Book, error) {

	checkAuthorExists, err := b.authorRepository.FindOneById(*body.AuthorID)
	if checkAuthorExists.ID == nil {
		fmt.Printf("author id : %v", *body.AuthorID)
		return nil, errors.New(constant.ResponseMsgAuthorDoesNotExist)
	}
	if err != nil {
		return nil, err
	}
	checkExist, err := b.bookRepository.FindSimilarBookByTitle(body.Title)
	if err != nil {
		return nil, err
	}
	for _, book := range checkExist {
		if body.Title == book.Title {
			return nil, errors.New(constant.ResponseMsgBookAlreadyExists)
		}
	}

	book, err := b.bookRepository.CreateBook(body)
	if err != nil {
		return nil, err
	}
	bookJson := utils.ConvertBookToJson(*book)
	return &bookJson, nil
}
