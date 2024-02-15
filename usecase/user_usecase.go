package usecase

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/constant"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/exception"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/utils"
)

type UserUseCase interface {
	GetUsers(title string) ([]dto.User, error)
	CreateUser(body dto.CreateUserBody) (*dto.User, error)
	Login(body dto.LoginBody) error
}

type userUseCaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserUseCaseImpl(userRepository repository.UserRepository) *userUseCaseImpl {
	return &userUseCaseImpl{
		userRepository: userRepository,
	}
}

func (u *userUseCaseImpl) GetUsers(name string) ([]dto.User, error) {

	usersJson := []dto.User{}
	var users []entity.User
	var err error
	if name == "" {
		users, err = u.userRepository.FindAll()

	} else {
		users, err = u.userRepository.FindSimilarUserByName(name)

	}
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		usersJson = append(usersJson, utils.ConvertUserToJson(user))
	}

	return usersJson, nil
}

func (u *userUseCaseImpl) Login(body dto.LoginBody) error {
	password, err := u.userRepository.FindUserPassword(body)
	if err != nil {
		return err
	}
	plainPassword, err := utils.CheckPassword(body.Password, []byte(password))
	if err != nil {
		return exception.NewErrorType(http.StatusUnauthorized, "Wrong Password")
	}
	if !plainPassword {
		return exception.NewErrorType(http.StatusUnauthorized, "Wrong Password")
	}
	return err
}

func (u *userUseCaseImpl) CreateUser(body dto.CreateUserBody) (*dto.User, error) {

	checkUserExist, err := u.userRepository.FindUserByEmail(body.Email)
	if checkUserExist.Email == body.Email {
		return nil, exception.NewErrorType(http.StatusBadRequest, constant.ResponseMsgUserAlreadyExists)
	}
	if err != nil {
		return nil, err
	}

	user, err := u.userRepository.CreateUser(body)
	if err != nil {
		return nil, err
	}
	userJson := utils.ConvertUserToJson(*user)
	return &userJson, nil

}
