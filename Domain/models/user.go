package models

type User struct {
	ID    string `json:"id",dynamodbav:"id",omitempty`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type CreateUser struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
