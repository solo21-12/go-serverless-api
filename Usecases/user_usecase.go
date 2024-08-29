package usecases

import (
	"github.com/solo21-12/go-server-less-API/Domain/interfaces"
	"github.com/solo21-12/go-server-less-API/Domain/models"
)

type userUseCase struct {
	emailService interfaces.EmailService
	repository   interfaces.UserRepository
}

func NewUserUseCase(emailService interfaces.EmailService,
	repository interfaces.UserRepository,
) interfaces.UserUseCase {
	return &userUseCase{
		emailService: emailService,
		repository:   repository,
	}
}

func (uc *userUseCase) GetUser(email string) (*models.User, *models.ErrorResponse) {
	return uc.repository.FetchUserID(email)
}

func (uc *userUseCase) CreateUser(newUser models.CreateUser) (*models.User, *models.ErrorResponse) {
	if user, err := uc.repository.FetchUserEmail(newUser.Email); err == nil && user != nil {
		return nil, err
	}

	if valid := uc.emailService.IsValidEmail(newUser.Email); !valid {
		return nil, models.BadRequest("Invalid Email address")
	}

	err := uc.repository.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	createdUser, nErr := uc.repository.FetchUserEmail(newUser.Email)
	if nErr != nil {
		return nil, nErr
	}

	return createdUser, nil
}

func (uc *userUseCase) UpdateUser(id string, updatedUser models.CreateUser) (*models.User, *models.ErrorResponse) {
	if _, err := uc.repository.FetchUserID(id); err != nil {
		return nil, err
	}

	err := uc.repository.UpdateUser(id, updatedUser)
	if err != nil {
		return nil, err
	}

	user, nErr := uc.repository.FetchUserID(id)
	if nErr != nil {
		return nil, nErr
	}

	return user, nil

}

func (uc *userUseCase) DeleteUser(id string) *models.ErrorResponse {

	if _, err := uc.repository.FetchUserID(id); err != nil {
		return err
	}

	err := uc.repository.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}
