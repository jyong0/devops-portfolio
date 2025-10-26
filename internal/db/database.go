package db

import (
	"context"
	"devops-portfolio/app/internal/config"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(cfg *config.Config) *pgxpool.Pool {
	dsn := "postgres://" + cfg.DBUser + ":" + cfg.DBPass +
		"@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	log.Println("Connected to PostgreSQL")
	return pool
}
