package postgres

import (
	"fmt"
	"github.com/jackc/pgx/v5" 
	"context"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func New(config Config) (*pgx.Conn, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.Username, config.Password, config.Host, config.Port, config.Database)
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return conn, nil
}