package api

import (
	"lambda-func/database"
)

type ApiHandler struct {
	dbStore *database.DynamoDBClient
}

func NewApiHandler(databaseStore *database.DynamoDBClient) *ApiHandler {
	return &ApiHandler{dbStore: databaseStore}
}
