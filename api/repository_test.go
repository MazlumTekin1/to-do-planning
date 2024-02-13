package api

import (
	"testing"
	"todo_planning/config"
	"todo_planning/database"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func Test_DistributeTasksController(t *testing.T) {}

func Test_AddTaskController(t *testing.T) {
	//load config info
	config := config.LoadConfig() //TODO: config ve db initialize işlemleri için error handling yapılmalı.
	//initialize database
	database, err := database.Connect(config.DatabaseURL)
	require.NoError(t, err)
	require.NotNil(t, database)

	dbProvider := DbProvider{Pool: database}
	require.NotNil(t, dbProvider)

	datas, err := dbProvider.GetDeveloperFromDatabase()
	require.NoError(t, err)
	require.NotNil(t, datas)
}

func Test_ApiHandler(t *testing.T) {
	//load config info
	config := config.LoadConfig() //TODO: config ve db initialize işlemleri için error handling yapılmalı.
	//initialize database
	database, err := database.Connect(config.DatabaseURL)
	require.NoError(t, err)
	require.NotNil(t, database)

	dbProvider := DbProvider{Pool: database}
	require.NotNil(t, dbProvider)

	provider := ApiHandler{
		DbProvider: dbProvider,
	}
	require.NotNil(t, provider)

	err = provider.DistributeTasksController(&fiber.Ctx{})
	require.Error(t, err, "panic: runtime error: invalid memory address or nil pointer dereference")

}
