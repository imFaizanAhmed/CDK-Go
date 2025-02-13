package api

import (
	"fmt"
	"lambda-func/database"
	types "lambda-func/type"
)

type ApiHandler struct {
	dbStore *database.DynamoDBClient
}

func NewApiHandler(databaseStore *database.DynamoDBClient) *ApiHandler {
	return &ApiHandler{dbStore: databaseStore}
}

func (api *ApiHandler) RegisterUserHandler(user types.RegisterUser) error {
	if user.Username == "" || user.Password == "" {
		return fmt.Errorf("username and password are required")
	}
	exists, err := api.dbStore.DoesUserExist(user.Username)
	if err != nil {
		return fmt.Errorf("error checking if user exists %w", err)
	}
	if exists {
		return fmt.Errorf("user already exists")
	}
	err = api.dbStore.InsertUser(user)
	if err != nil {
		return fmt.Errorf("error inserting user %w", err)
	}
	return nil
}
