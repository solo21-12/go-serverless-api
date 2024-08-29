package routers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/solo21-12/go-server-less-API/Delivery/controllers"
	"github.com/solo21-12/go-server-less-API/config"
)

func NewUserRouter(env *config.Env, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	handlers := controllers.NewUserControllers(env)

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
