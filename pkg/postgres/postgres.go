package postgres

import (
	"context"
	"fmt"
	"github.com/Hellsker/Todo-List/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type DbConfig struct {
	Username string
	Password string
	Hostname string
	Port     string
	DBName   string
}

func NewConfig(cfg *config.Config) *DbConfig {
	return &DbConfig{
		Username: cfg.Database.Username,
		Password: cfg.Database.Password,
		Hostname: cfg.Database.Host,
		Port:     cfg.Database.Port,
		DBName:   cfg.Database.DBName,
	}
}

// DSN will get datasource name of the database configuration
func (c *DbConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		c.Username, c.Password, c.Hostname, c.Port, c.DBName)
}

// Postgres is pgxpool.Poll wrapper
type Postgres struct {
	pool *pgxpool.Pool
}

// Singleton pattern
var (
	pgInstance *Postgres
	pgOnce     sync.Once
)

// New
func New(dbConfig *DbConfig) *Postgres {
	pgOnce.Do(func() {
		cfg, err := pgxpool.ParseConfig(dbConfig.DSN())
		if err != nil {
			return
		}
		pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
		if err != nil {
			return
		}
		pgInstance = &Postgres{pool: pool}
	})
	return pgInstance
}
func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.pool.Ping(ctx)
}
func (pg *Postgres) Close() {
	if pg.pool != nil {
		pg.pool.Close()
	}
}
func (pg *Postgres) GetPool() *pgxpool.Pool {
	return pg.pool
}
