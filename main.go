package main

import (
	"fmt"
	"todo_planning/api"
	"todo_planning/config"
	"todo_planning/database"
	"todo_planning/handlers"
)

func main() {
	//load config info
	config := config.LoadConfig()
	//initialize database
	conn := database.Initialize()
	defer conn.Close()

	//initialize provider
	provider := api.DbProvider{Pool: conn}

	// Truncate the tasks table
	provider.TruncateTasksFromDatabase()

	provider1 := api.Provider1{Db: &provider}
	provider1Tasks, err := provider1.GetTasks(config.Provider1URL) //TODO: Burada Provider1 ve Provedir2 için tek bir yapı gerekecek.
	if err != nil {
		fmt.Println("Error getting tasks from provider 1: ", err)
	}
	// provider2 := api.Provider2{Db: &provider}
	// provider2Tasks, err := provider2.GetTasks(config.Provider2URL)
	// if err != nil {
	// 	fmt.Println("Error getting tasks from provider 2: ", err)
	// }
	dev := []handlers.Developer{
		{
			ID:     "muh",
			Cap:    1,
			Rating: 1,
		},
		{
			ID:     "asd",
			Cap:    1,
			Rating: 2,
		},
		{
			ID:     "dfg",
			Cap:    1,
			Rating: 3,
		},
		{
			ID:     "qwe",
			Cap:    1,
			Rating: 4,
		},
		{
			ID:     "zxc",
			Cap:    1,
			Rating: 5,
		},
	}

	handlers.WeeklyWorkPlan(provider1Tasks, dev, 45)
}
