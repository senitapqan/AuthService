package server

import (
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	dbInstance   *sqlx.DB
	once sync.Once
)

type Config struct {
	Host string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func connectionStringBuilder(cfg Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
}

func NewPostgresConnection(cfg Config) (*sqlx.DB, error) {
	var err error
	once.Do(func() {
		dbInstance, err = sqlx.Open("postgres", connectionStringBuilder(cfg))
		if err != nil {
			return
		}
		err = dbInstance.Ping()
		if err != nil {
			return
		}
	})
	if err != nil {
		return nil, err
	}
	return dbInstance, nil
}
