package main

import (
	"log"
	"todo_planning/api"
	"todo_planning/config"
	"todo_planning/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	//load config info
	config := config.LoadConfig() //TODO: config ve db initialize işlemleri için error handling yapılmalı.
	//initialize database
	database, err := database.Connect(config.DatabaseURL)
	if err != nil {
		log.Fatal("Database Connection Can't Estabilished, Error:", err)
	} else {
		log.Println("Database Connection Estabilished")
	}
	//TODO logger yapılmalı
	//initialize provider

	dbProvider := api.DbProvider{Pool: database}

	provider := api.ApiHandler{
		DbProvider: dbProvider,
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	app.Route("/tasks", provider.Router)
	app.Listen(":3000")
}
