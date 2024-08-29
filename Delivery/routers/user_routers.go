package routers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/solo21-12/go-server-less-API/Delivery/controllers"
	infrastructure "github.com/solo21-12/go-server-less-API/Infrastructure"
	repository "github.com/solo21-12/go-server-less-API/Repository"
	usecases "github.com/solo21-12/go-server-less-API/Usecases"
	"github.com/solo21-12/go-server-less-API/config"
)

func NewUserRouter(env *config.Env, req events.APIGatewayProxyRequest, dynamoClient *dynamodb.DynamoDB) (events.APIGatewayProxyResponse, error) {
	userRepo := repository.NewUserRepository(dynamoClient)
	emailService := infrastructure.NewEmailService(*env)
	useCase := usecases.NewUserUseCase(emailService, userRepo)
	handlers := controllers.NewUserControllers(env, useCase)

	switch req.HTTPMethod {
	case "GET":
		handlers.GetUser(&req)
	case "POST":
		handlers.CreateUser(&req)
	case "PUT":
		handlers.UpdateUser(&req)
	case "DELETE":
		handlers.DeleteUser(&req)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 405,
		Body:       "Method not allowed",
	}, nil
}
