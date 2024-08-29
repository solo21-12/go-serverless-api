package interfaces

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/solo21-12/go-server-less-API/Domain/models"
)

type UserControllers interface {
	GetUser(events *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error)
	CreateUser(events *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error)
	UpdateUser(events *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error)
	DeleteUser(events *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error)
}

type UserUseCase interface {
	GetUser(Id string) (*models.User, *models.ErrorResponse)
	CreateUser(newUser models.CreateUser) (*models.User, *models.ErrorResponse)
	UpdateUser(id string, updatedUser models.CreateUser) (*models.User, *models.ErrorResponse)
	DeleteUser(id string) *models.ErrorResponse
}

type UserRepository interface {
	FetchUserEmail(email string) (*models.User, *models.ErrorResponse)
	FetchUserID(id string) (*models.User, *models.ErrorResponse)
	Fetchusers() ([]models.User, *models.ErrorResponse)
	CreateUser(user models.CreateUser) *models.ErrorResponse
	UpdateUser(id string, user models.CreateUser) *models.ErrorResponse
	DeleteUser(id string) *models.ErrorResponse
}
