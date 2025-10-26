package service

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserService struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewUserService(db *pgxpool.Pool, rdb *redis.Client) *UserService {
	return &UserService{db: db, rdb: rdb}
}

func (s *UserService) CreateUser(ctx context.Context, name, email string) error {
	_, err := s.db.Exec(ctx, `INSERT INTO users (name, email) VALUES ($1, $2)`, name, email)
	return err
}

func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
	cacheData, err := s.rdb.Get(ctx, "user:"+id).Result()
	if err == nil {
		log.Println("[CACHE HIT]")
		var u User
		json.Unmarshal([]byte(cacheData), &u)
		return &u, nil
	}

	log.Println("[CACHE MISS] Fetching from DB...")
	row := s.db.QueryRow(ctx, `SELECT id, name, email FROM users WHERE id=$1`, id)
	var u User
	err = row.Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return nil, err
	}

	jsonData, _ := json.Marshal(u)
	s.rdb.Set(ctx, "user:"+id, string(jsonData), 0)

	return &u, nil
}
