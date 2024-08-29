package config

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewDynamoClient(env *Env) *dynamodb.DynamoDB {
	region := env.AWS_REGION
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		log.Fatalf("Error creating AWS session: %v", err)
	}

	return dynamodb.New(awsSession)
}
