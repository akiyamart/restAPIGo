package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	todo "github.com/akiyamart/restAPIGo"
	"github.com/akiyamart/restAPIGo/pkg/handler"
	"github.com/akiyamart/restAPIGo/pkg/repository"
	"github.com/akiyamart/restAPIGo/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
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
	go func () { 
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil { 
			log.Fatalf("error occured while running http sever: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit

	logrus.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil { 
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil { 
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func InitConfig() error { 
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}