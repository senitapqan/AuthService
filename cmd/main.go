package main

import (
	"goAuthService/internal/repository"
	"goAuthService/internal/service"
	"goAuthService/internal/handler"
	
	"goAuthService/server"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if err := initConfig(); err != nil {
		log.Fatal().Err(err).Msg("some error with initializiing")
			
	}

	db, err := server.NewPostgresConnection(server.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatal().Err(err).Msg("error with data base")
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(server.Server)

	if err := srv.RunServer(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatal().Err(err).Msg("error with run server")
	}
}

func initConfig() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("configs")
	return viper.ReadInConfig()
}