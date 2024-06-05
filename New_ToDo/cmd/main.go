package main

import (
	newtodo "exapmle/todo_app"
	"exapmle/todo_app/pkg/handler"
	"exapmle/todo_app/pkg/repository"
	"exapmle/todo_app/pkg/service"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load.env file:  %s", err.Error())
	}

	db, err := repository.NewPostgresBD(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db:  %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(newtodo.Server)
	if err := srv.Run(viper.GetString("8800"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
