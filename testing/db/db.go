// Package db provides helper and utility methods for testing with a database.
package db

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // needed to load driver
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/rs/zerolog"
)

const (
	postgresDatabaseEnv = "TESTING_POSTGRES_DATABASE"
	postgresUsernameEnv = "TESTING_POSTGRES_MIGRATION_USERNAME"
	//nolint: gosec // credentials only used for testing
	postgresPasswordEnv = "TESTING_POSTGRES_MIGRATION_PASSWORD"
	postgresHostEnv     = "TESTING_POSTGRES_HOST"
)

// IntegrationDB is the struct that holds all the necessary dependencies for us to run db tests
type IntegrationDB struct {
	logger          zerolog.Logger
	DB              *sqlx.DB
	dockerPool      *dockertest.Pool
	dockerResource  *dockertest.Resource
	DBConnectionURL string
}

// NewIntegrationDB creates all the resources necessary for us to run database integration tests
func NewIntegrationDB() (*IntegrationDB, error) {
	var db *sqlx.DB
	var pool *dockertest.Pool
	var resource *dockertest.Resource
	var DBConnectionURL string
	var err error
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	if x := os.Getenv(postgresDatabaseEnv); x != "" {
		logger.Info().Msg("connecting to existing postgres instance - starting postgres via Go process")
		db, DBConnectionURL, err = ConnectPostgres()
		if err != nil {
			return nil, fmt.Errorf("while connecting to postgres using: %s error: %w", DBConnectionURL, err)
		}
	} else {
		logger.Info().Msg("starting postgres using ory-dockertest")
		pool, resource, db, DBConnectionURL, err = StartAndConnectPostgres()
		if err != nil {
			logger.Err(err).Msg("failed to Start and connect to postgres")
			return nil, err
		}
	}

	return &IntegrationDB{
		logger:          logger,
		DB:              db,
		dockerPool:      pool,
		dockerResource:  resource,
		DBConnectionURL: DBConnectionURL,
	}, nil
}

// ConnectPostgres connects to an already running postgres instance.
func ConnectPostgres() (*sqlx.DB, string, error) {
	user := os.Getenv(postgresUsernameEnv)
	password := os.Getenv(postgresPasswordEnv)
	host := os.Getenv(postgresHostEnv)
	dbName := os.Getenv(postgresDatabaseEnv)

	// Build the connection URL.
	connectionURL := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(user, password),
		Host:     host,
		Path:     dbName,
		RawQuery: "sslmode=disable",
	}

	// attempt to connect with simple retries
	var db *sqlx.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = sqlx.Connect("postgres", connectionURL.String())
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		break
	}
	if err != nil {
		return nil, connectionURL.String(), fmt.Errorf("error connecting to postgres: %w", err)
	}

	return db, connectionURL.String(), nil
}

// StartAndConnectPostgres starts a postgres docker container, runs migrations, and connects to it.
func StartAndConnectPostgres() (*dockertest.Pool, *dockertest.Resource, *sqlx.DB, string, error) {

	//nolint: goconst // credentials only used for testing
	user := "test"
	password := "test"
	dbName := "test"

	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, nil, nil, "", fmt.Errorf("could not connect to docker: %w", err)
	}

	if err := pool.Client.Ping(); err != nil {
		return nil, nil, nil, "", fmt.Errorf("could not connect to Docker at Endpoint: %s. Error is: %w", pool.Client.Endpoint(), err)
	}

	options := &dockertest.RunOptions{
		Name:       randomContainerName(),
		Repository: "postgres",
		Tag:        "12",
		Env: []string{
			fmt.Sprintf("POSTGRES_USER=%s", user),
			fmt.Sprintf("POSTGRES_PASSWORD=%s", password),
			fmt.Sprintf("POSTGRES_DB=%s", dbName),
		},
	}
	resource, err := pool.RunWithOptions(options, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})

	if err != nil {
		return nil, nil, nil, "", fmt.Errorf("could not start resource: %w", err)
	}

	// Build the connection URL.
	connectionURL := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(user, password),
		Host:     resource.GetHostPort("5432/tcp"),
		Path:     dbName,
		RawQuery: "sslmode=disable",
	}

	var db *sqlx.DB
	var dbErr error
	if err = pool.Retry(func() error {
		db, dbErr = sqlx.Connect("postgres", connectionURL.String())
		if dbErr != nil {
			return fmt.Errorf("open db: %w", dbErr)
		}

		return nil
	}); err != nil {
		return nil, nil, nil, "", fmt.Errorf("could not connect to postgres container: %w", dbErr)
	}

	return pool, resource, db, connectionURL.String(), nil
}

// randomContainerName returns a random docker container name. All containers start with application-test- for easy cleanup.
func randomContainerName() string {
	b := make([]byte, 4)
	if _, err := rand.Read(b); err != nil {
		panic(err) // should never panic, unless something is very wrong
	}
	return "service-test-" + hex.EncodeToString(b)
}

// CleanUpDB kills the database connection when test is done
func (d IntegrationDB) CleanUpDB() error {
	if x := os.Getenv(postgresDatabaseEnv); x != "" {
		// clean up resources for an existing postgres instance
		d.logger.Info().Msg("closing dbx connections for tests")
		if err := d.DB.Close(); err != nil {
			return fmt.Errorf("while closing dbx for tests: %w", err)
		}
	} else {
		// kill the postgres container when test is done
		if err := d.dockerPool.Purge(d.dockerResource); err != nil {
			return fmt.Errorf("could not purge resource: %w", err)
		}

	}

	return nil
}
