package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/solo21-12/go-server-less-API/Delivery/routers"
)

func main() {
	lambda.Start(routers.Handler)
}
