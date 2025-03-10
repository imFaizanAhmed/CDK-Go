package database

import (
	types "lambda-func/type"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	TABLE_NAME = "userTable"
)

type DynamoDBClient struct {
	databaseStore *dynamodb.DynamoDB
}

func NewDynamoDBClient() *DynamoDBClient {

	dbSession := session.Must(session.NewSession())
	db := dynamodb.New(dbSession)

	return &DynamoDBClient{databaseStore: db}
}

func (u *DynamoDBClient) DoesUserExist(username string) (bool, error) {
	result, err := u.databaseStore.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {S: aws.String(username)},
		},
	})
	if err != nil {
		return false, err
	}
	return result.Item != nil, nil
}

func (u *DynamoDBClient) InsertUser(user types.RegisterUser) error {
	Item := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item: map[string]*dynamodb.AttributeValue{
			"username": {S: aws.String(user.Username)},
			"password": {S: aws.String(user.Password)},
		},
	}

	_, err := u.databaseStore.PutItem(Item)
	if err != nil {
		return err
	}
	return nil
}
