package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/solo21-12/go-server-less-API/Domain/interfaces"
	"github.com/solo21-12/go-server-less-API/Domain/models"
	"github.com/solo21-12/go-server-less-API/config"
)

type UserControllers struct {
	env     *config.Env
	useCase interfaces.UserUseCase
}

func NewUserControllers(
	env *config.Env,
	usecase interfaces.UserUseCase,
) interfaces.UserControllers {
	return &UserControllers{env: env, useCase: usecase}
}

func (uc *UserControllers) GetUser(events *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := events.PathParameters["id"]

	user, err := uc.useCase.GetUser(id)
	if err != nil {
		return interfaces.ApiResponse(err.Code, models.ErrorBody{
			ErrorMsg: &err.Message,
		})
	}
	return interfaces.ApiResponse(200, user)
}

func (uc *UserControllers) CreateUser(events *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	newUser := models.CreateUser{}

	if err := json.Unmarshal([]byte(events.Body), &newUser); err != nil {
		errMsg := "Invalid request body"
		return interfaces.ApiResponse(http.StatusBadRequest, models.ErrorBody{

			ErrorMsg: &errMsg,
		})
	}

	createdUser, cErr := uc.useCase.CreateUser(newUser)

	if cErr != nil {
		return interfaces.ApiResponse(cErr.Code, models.ErrorBody{
			ErrorMsg: &cErr.Message,
		})
	}

	return interfaces.ApiResponse(http.StatusCreated, createdUser)
}

func (uc *UserControllers) UpdateUser(events *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	updatedUser := models.CreateUser{}
	id := events.PathParameters["id"]

	if err := json.Unmarshal([]byte(events.Body), &updatedUser); err != nil {
		errMsg := "Invalid request body"
		return interfaces.ApiResponse(http.StatusBadRequest, models.ErrorBody{

			ErrorMsg: &errMsg,
		})
	}

	createdUser, cErr := uc.useCase.UpdateUser(id, updatedUser)

	if cErr != nil {
		return interfaces.ApiResponse(cErr.Code, models.ErrorBody{
			ErrorMsg: &cErr.Message,
		})
	}

	return interfaces.ApiResponse(http.StatusCreated, createdUser)
}

func (uc *UserControllers) DeleteUser(events *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	id := events.PathParameters["id"]

	err := uc.useCase.DeleteUser(id)
	if err != nil {
		return interfaces.ApiResponse(err.Code, models.ErrorBody{
			ErrorMsg: &err.Message,
		})

	}

	return interfaces.ApiResponse(http.StatusNoContent, nil)
}
