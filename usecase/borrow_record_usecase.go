package usecase

import (
	"context"
	"errors"
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/constant"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/database"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/exception"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/utils"
)

type BorrowRecordUseCase interface {
	NewBorrowRecord(ctx context.Context, body dto.CreateBorrowRecordBody, id int) (*dto.BorrowRecord, error)
	ReturnBorrowedBook(ctx context.Context, id int) (*dto.BorrowRecord, error)
}

type borrowRecordUseCaseImpl struct {
	borrowRecordRepository repository.BorrowRecordRepository
	bookRepository         repository.BookRepository
	userRepository         repository.UserRepository
	transactor             database.Transactor
}

func NewBorrowRecordUseCaseImpl(
	borrowRecordRepository repository.BorrowRecordRepository,
	bookRepository repository.BookRepository,
	userRepository repository.UserRepository,
	transactor database.Transactor,
) *borrowRecordUseCaseImpl {
	return &borrowRecordUseCaseImpl{
		borrowRecordRepository: borrowRecordRepository,
		bookRepository:         bookRepository,
		userRepository:         userRepository,
		transactor:             transactor,
	}
}

func (r *borrowRecordUseCaseImpl) NewBorrowRecord(ctx context.Context, body dto.CreateBorrowRecordBody, id int) (*dto.BorrowRecord, error) {

	var recordJson dto.BorrowRecord
	err := r.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {

		checkBookExists, err := r.bookRepository.FindOneById(txCtx, body.BookID)
		if checkBookExists == nil {
			return errors.New(constant.ResponseMsgBookDoesNotExist)
		}
		if checkBookExists.Quantity < 1 {
			return errors.New(constant.ResponseMsgBookDoesNotExist)
		}
		if err != nil {
			return err
		}

		decreasedBook, err := r.bookRepository.DecreaseBookQuantity(txCtx, body.BookID)
		if decreasedBook == nil {
			return errors.New(constant.ResponseMsgBookDoesNotExist)
		}
		if err != nil {
			return err
		}
		bRecord, err := r.borrowRecordRepository.CreateBorrowRecord(txCtx, body, int64(id))
		if err != nil {
			return err
		}
		recordJson = utils.ConvertBorrowRecordToJson(*bRecord)
		return nil
	})

	if err != nil {
		return nil, exception.NewErrorType(http.StatusBadRequest, err.Error())
	}
	return &recordJson, nil
}

func (r *borrowRecordUseCaseImpl) ReturnBorrowedBook(ctx context.Context, id int) (*dto.BorrowRecord, error) {

	checkRecordExists, err := r.borrowRecordRepository.FindOneById(ctx, int64(id))
	if checkRecordExists.ID == 0 {
		return nil, exception.NewErrorType(
			http.StatusPreconditionFailed, constant.ResponseMsgRecordDoesNotExist)
	}
	if checkRecordExists.Status == "returned" {
		return nil, exception.NewErrorType(
			http.StatusPreconditionFailed, constant.ResponseMsgBookAlreadyReturned)
	}
	if err != nil {
		return nil, err
	}

	returnedBook, err := r.bookRepository.IncreaseBookQuantity(ctx, checkRecordExists.BookID)
	if returnedBook == nil {
		return nil, errors.New(constant.ResponseMsgBookDoesNotExist)
	}
	if err != nil {
		return nil, err
	}
	updatedRecord, err := r.borrowRecordRepository.UpdateRecordReturnBook(ctx, int64(id))
	if err != nil {
		return nil, err
	}
	recordJson := utils.ConvertBorrowRecordToJson(*updatedRecord)

	return &recordJson, nil
}
