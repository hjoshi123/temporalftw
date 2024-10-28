package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/hjoshi123/temporal-loan-app/internal/config"
	"github.com/hjoshi123/temporal-loan-app/internal/logging"
	"sync"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	db   *sql.DB
	once sync.Once
)

func Connect(ctx context.Context) *sql.DB {
	once.Do(func() {
		localDB, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			config.Spec.DBUser, config.Spec.DBPassword, config.Spec.DBHost, config.Spec.DBPort, config.Spec.DBName))
		if err != nil {
			logging.FromContext(ctx).Panicw("failed to connect to database", "error", err)
		}

		if db == nil {
			db = localDB
		}
	})

	boil.SetDB(db)
	return db
}
