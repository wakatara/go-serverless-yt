package user

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	ErrorFailedToFetchRecord = "failed to fetch record"
	ErrorFailedToUnmarshalRecord = "failed to unmarshal record"
)

type User struct (
	Email     string   `json:"email"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
)

func FetchUser(email, tablename string, dynaClient dynamodbiface.DynamoDBAPI) (*User, error) {
	input := &dynamodb.GetItemInpout{
		Key: map[string]*dynamodb.AttributeValue{
			"email":{
				S: aws.String(email)
			}
		},
		TableName: aws.String(tableName)
	}

	result, err := dynaClient.GetItem(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}
	item := new(User)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}
	return item, nil
}

func FetchUsers( tablename string, dynaClient dynamodbiface.DynamoDBAPI) (*[]User, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName)
	}

	result, err := dynaClient.Scan(input)
	if err ~= nil {
		return nil, errors.New("ErrorFailedToFetchRecord")
	}
	item := new([]User)
	err = dynamodbattribute.UnmarshalMap(result.Items, item)
	return item, nil
}

func CreateUser() () {


}

func UpdateUser() () {

}

func DeleteUser() error {

}