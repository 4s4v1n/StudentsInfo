package postgres_db

import (
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

const (
	maxOpenConns    = 50
	maxIdleConns    = 25
	connMaxLifetime = time.Duration(30)
)

type PostgresDb struct {
	DB  *sqlx.DB
	cfg PostgresConfig
}

type PostgresConfig struct {
	DSN             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func NewPgx(opts ...Option) (*PostgresDb, error) {
	postgresDb := &PostgresDb{
		cfg: PostgresConfig{
			MaxOpenConns:    maxOpenConns,
			MaxIdleConns:    maxIdleConns,
			ConnMaxLifetime: connMaxLifetime,
		},
	}
	for _, opt := range opts {
		opt(postgresDb)
	}

	db, err := sqlx.Open("pgx", postgresDb.cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("open database connection: %w", err)
	}

	db.SetMaxOpenConns(postgresDb.cfg.MaxOpenConns)
	db.SetMaxIdleConns(postgresDb.cfg.MaxIdleConns)
	db.SetConnMaxLifetime(postgresDb.cfg.ConnMaxLifetime)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}

	postgresDb.DB = db
	return postgresDb, nil
}

func (p *PostgresDb) Close() {
	if err := p.DB.Close(); err != nil {
		log.Error().Msgf("Can't close database client: %v", err)
	}
}
