package main

import (
	"log"
	"os"

	todo "github.com/akiyamart/restAPIGo"
	"github.com/akiyamart/restAPIGo/pkg/handler"
	"github.com/akiyamart/restAPIGo/pkg/repository"
	"github.com/akiyamart/restAPIGo/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := InitConfig(); err != nil { 
		log.Fatalf("error initializing config: %s", err.Error())
	}

	if err:= godotenv.Load(); err != nil { 
		log.Fatalf("eror loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil { 
		log.Fatalf("failed to initialize database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil { 
		log.Fatalf("error occured while running http sever: %s", err.Error())
	}
}

func InitConfig() error { 
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}