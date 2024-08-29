package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/solo21-12/go-server-less-API/Domain/interfaces"
	"github.com/solo21-12/go-server-less-API/Domain/models"
)

type userRepository struct {
	tableName  string
	dynaClient dynamodbiface.DynamoDBAPI
}

func NewUserRepository(dynaClient dynamodbiface.DynamoDBAPI) interfaces.UserRepository {
	return &userRepository{
		tableName:  "users",
		dynaClient: dynaClient,
	}
}

func (r *userRepository) FetchUser(key string, value string) (*models.User, *models.ErrorResponse) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			key: {
				S: aws.String(value),
			},
		},
	}

	result, err := r.dynaClient.GetItem(input)
	if err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	user := new(models.User)
	if err := dynamodbattribute.UnmarshalMap(result.Item, &user); err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	return user, nil
}

func (r *userRepository) FetchUserEmail(email string) (*models.User, *models.ErrorResponse) {
	return r.FetchUser("email", email)
}

func (r *userRepository) FetchUserID(id string) (*models.User, *models.ErrorResponse) {
	return r.FetchUser("id", id)
}

func (r *userRepository) Fetchusers() ([]models.User, *models.ErrorResponse) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(r.tableName),
	}

	result, err := r.dynaClient.Scan(input)
	if err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	users := make([]models.User, 0)
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &users); err != nil {
		return nil, models.InternalServerError(err.Error())
	}

	return users, nil
}

func (r *userRepository) CreateUser(user models.CreateUser) *models.ErrorResponse {

	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return models.InternalServerError(err.Error())
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      av,
	}

	_, err = r.dynaClient.PutItem(input)
	if err != nil {
		return models.InternalServerError(err.Error())
	}

	return nil

}

func (r *userRepository) UpdateUser(id string, user models.CreateUser) *models.ErrorResponse {
	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return models.InternalServerError(err.Error())
	}

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":u": {
				M: av,
			},
		},
		UpdateExpression: aws.String("SET user = :u"),
		ReturnValues:     aws.String("ALL_NEW"),
	}

	_, err = r.dynaClient.UpdateItem(input)
	if err != nil {
		return models.InternalServerError(err.Error())
	}

	return nil
}

func (r *userRepository) DeleteUser(id string) *models.ErrorResponse {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}

	_, err := r.dynaClient.DeleteItem(input)
	if err != nil {
		return models.InternalServerError(err.Error())
	}

	return nil
}
