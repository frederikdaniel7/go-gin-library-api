package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/utils"
)

type BookUseCase interface {
	GetBooks(title string) ([]dto.Book, error)
}

type bookUseCaseImpl struct {
	bookRepository repository.BookRepository
}

func NewBookUseCaseImpl(bookRepository repository.BookRepository) *bookUseCaseImpl {
	return &bookUseCaseImpl{
		bookRepository: bookRepository,
	}
}

func (b *bookUseCaseImpl) GetBooks(title string) ([]dto.Book, error) {

	booksJson := []dto.Book{}
	var books []entity.Book
	var err error
	if title == "" {
		books, err = b.bookRepository.FindAll()

	} else {
		books, err = b.bookRepository.FindOneBookByTitle(title)

	}
	if err != nil {
		return nil, err
	}
	for _, book := range books {
		booksJson = append(booksJson, utils.ConvertBookToJson(book))
	}

	return booksJson, nil

}
