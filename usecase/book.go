package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/repository"
)

type BookUseCase interface{
	GetAll()([]entity.Book, error)
}

type bookUseCaseImpl struct {
	bookRepository repository.BookRepository
}

func NewBookUseCaseImpl(bookRepository repository.BookRepository) *bookUseCaseImpl {
	return &bookUseCaseImpl{
		bookRepository: bookRepository,
	}
}

func (b *bookUseCaseImpl) GetAll()([]entity.Book, error) {
    books, err := b.bookRepository.FindAll()
	if err != nil{
		return nil, err
	}
	return books, nil

}