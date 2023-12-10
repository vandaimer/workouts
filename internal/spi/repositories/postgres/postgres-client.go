package postgres

import (
	"database/sql"
	"net/url"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/vandaimer/workouts/config"
)

type BunPostgresDatabaseClient struct {
	DB     *bun.DB
	config *config.DbConfig
}

func (client *BunPostgresDatabaseClient) getPostgresURL() string {
	return "postgres://" + client.config.DB_DBName + ":" + url.QueryEscape(client.config.DB_Password) + "@" + client.config.DB_Host + ":" + client.config.DB_Port + "/" + client.config.DB_DBName + "?sslmode=disable"
}

func NewBunPostgresDatabaseClient(config *config.DbConfig) *BunPostgresDatabaseClient {
	client := &BunPostgresDatabaseClient{}
	client.config = config
	client.Connect()
	log.Info().Msg("Database client initialized.")
	return client
}

func (client *BunPostgresDatabaseClient) Connect() error {
	connectionString := client.getPostgresURL()
	log.Info().Msgf("Connecting to database: %s", connectionString)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connectionString)))
	client.DB = bun.NewDB(sqldb, pgdialect.New())
	err := client.DB.Ping()
	return err
}