package routers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/solo21-12/go-server-less-API/config"
)

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	env := config.NewEnv()
	dynamoClient := config.NewDynamoClient(env)

	switch req.Path {
	case "/user":
		return NewUserRouter(env, req, dynamoClient)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 405,
		Body:       "Method not allowed",
	}, nil
}
