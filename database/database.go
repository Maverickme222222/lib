// Package database ...
package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.elastic.co/apm/module/apmsql/v2"
	_ "go.elastic.co/apm/module/apmsql/v2/pq" // register postgres driver
)

// Config is the config that holds database connection details
type Config struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
}

// ConnectDBWithTracing creates a database with tracing capabilities that is used within the service
func ConnectDBWithTracing(config Config) (*sqlx.DB, error) {
	db, err := apmsql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DatabaseName))
	if err != nil {
		return nil, err
	}

	return sqlx.NewDb(db, "postgres"), nil
}
