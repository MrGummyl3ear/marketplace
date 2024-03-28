package main

import (
	"marketplace"
	"marketplace/pkg/handler"
	"marketplace/pkg/repository"
	"marketplace/pkg/service"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_Host"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_Username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DBName"),
		SSLMode:  os.Getenv("SSLMode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s ", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(marketplace.Server)

	if err := srv.Run(os.Getenv("Server_PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server:  %s", err.Error())
	}
}
