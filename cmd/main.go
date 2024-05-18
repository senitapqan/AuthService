package main

import (
	"goAuthService/internal/handler"
	"goAuthService/internal/repository"
	"goAuthService/internal/service"

	"goAuthService/server"

	"github.com/joho/godotenv"
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
		Password: viper.GetString("db.password"),
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
	if err := godotenv.Load(); err != nil {
		return err
	}
	viper.AddConfigPath("./config")
	viper.SetConfigName("configs")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	viper.BindEnv("db.password", "DB_PASSWORD")
	return nil
}
