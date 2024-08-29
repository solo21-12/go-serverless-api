package controllers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/solo21-12/go-server-less-API/Domain/interfaces"
	"github.com/solo21-12/go-server-less-API/config"
)

type UserControllers struct {
	env *config.Env
}

func NewUserControllers(env *config.Env) interfaces.UserControllers {
	return &UserControllers{env}
}

func (uc *UserControllers) GetUser(events *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return nil, nil
}

func (uc *UserControllers) CreateUser(events *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return nil, nil

}

func (uc *UserControllers) UpdateUser(events *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return nil, nil
}

func (uc *UserControllers) DeleteUser(events *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return nil, nil
}
