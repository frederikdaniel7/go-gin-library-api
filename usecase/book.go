package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/utils"
)

type BookUseCase interface {
	GetAll() ([]dto.Book, error)
	GetBookByTitle(title string) ([]dto.Book, error)
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

func (b *bookUseCaseImpl) GetAll() ([]dto.Book, error) {
	books, err := b.bookRepository.FindAll()
	if err != nil {
		return nil, err
	}

	booksJson := []dto.Book{}
	for _, book := range books {
		booksJson = append(booksJson, utils.ConvertBookToJson(book))
	}

	return booksJson, nil

}

func (b *bookUseCaseImpl) GetBookByTitle(title string) ([]dto.Book, error) {
	books, err := b.bookRepository.FindOneBookByTitle(title)
	if err != nil {
		return nil, err
	}
	booksJson := []dto.Book{}
	for _, book := range books {
		booksJson = append(booksJson, utils.ConvertBookToJson(book))
	}

	return booksJson, nil

}

func (b *bookUseCaseImpl) GetBooks(title string) ([]dto.Book, error) {

	booksJson := []dto.Book{}

	if title == "" {
		books, err := b.bookRepository.FindAll()
		for _, book := range books {
			booksJson = append(booksJson, utils.ConvertBookToJson(book))
		}
		if err != nil {
			return nil, err
		}
	} else {
		books, err := b.bookRepository.FindOneBookByTitle(title)
		for _, book := range books {
			booksJson = append(booksJson, utils.ConvertBookToJson(book))
		}
		if err != nil {
			return nil, err
		}
	}

	return booksJson, nil

}
