package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/entity"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-03/frederik-hutabarat/exercise-library-api/utils"
)

type UserUseCase interface {
	GetUsers(title string) ([]dto.User, error)
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
