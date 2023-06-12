package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func ConnPostgres(host, user string, port int, password, dbname string) (*DB, error) {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/%s ",
		user, password, host, port, dbname)

	pool, err := pgxpool.New(context.Background(), psqlInfo)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err = pool.Ping(ctx); err != nil {
		return nil, err
	}

	DB := &DB{
		Pool: pool,
	}
	return DB, nil
}
