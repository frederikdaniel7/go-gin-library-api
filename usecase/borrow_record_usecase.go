package usecase

import (
	"errors"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/constant"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/utils"
)

type BorrowRecordUseCase interface {
	NewBorrowRecord(body dto.CreateBorrowRecordBody) (*entity.BorrowRecord, error)
}

type borrowRecordUseCaseImpl struct {
	borrowRecordRepository repository.BorrowRecordRepository
	bookRepository         repository.BookRepository
	userRepository         repository.UserRepository
}

func NewBorrowRecordUseCaseImpl(
	borrowRecordRepository repository.BorrowRecordRepository,
	bookRepository repository.BookRepository,
	userRepository repository.UserRepository,
) *borrowRecordUseCaseImpl {
	return &borrowRecordUseCaseImpl{
		borrowRecordRepository: borrowRecordRepository,
		bookRepository:         bookRepository,
		userRepository:         userRepository,
	}
}

func (r *borrowRecordUseCaseImpl) NewBorrowRecord(body dto.CreateBorrowRecordBody) (*dto.BorrowRecord, error) {
	checkUserExists, err := r.userRepository.FindUserById(body.UserID)
	if checkUserExists == nil {
		return nil, errors.New(constant.ResponseMsgAuthorDoesNotExist)
	}
	if err != nil {
		return nil, err
	}
	checkBookExists, err := r.bookRepository.FindOneById(body.BookID)
	if checkBookExists == nil {
		return nil, errors.New(constant.ResponseMsgAuthorDoesNotExist)
	}
	if err != nil {
		return nil, err
	}
	bRecord, err := r.borrowRecordRepository.CreateBorrowRecord(body)
	if err != nil {
		return nil, err
	}
	recordJson := utils.ConvertBorrowRecordToJson(*bRecord)
	return &recordJson, nil

}
