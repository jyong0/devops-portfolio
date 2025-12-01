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
	Age   int    `json:"age,omitempty"`
}

type UserService struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewUserService(db *pgxpool.Pool, rdb *redis.Client) *UserService {
	return &UserService{db: db, rdb: rdb}
}

func (s *UserService) CreateUser(ctx context.Context, name, email string, age int) error {
	_, err := s.db.Exec(ctx,
		`INSERT INTO users (name, email, age) VALUES ($1, $2, $3)`,
		name, email, age,
	)
	return err
}

func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
	// 🔹 CACHE CHECK
	cacheData, err := s.rdb.Get(ctx, "user:"+id).Result()
	if err == nil {
		log.Println("[CACHE HIT]")
		var u User
		json.Unmarshal([]byte(cacheData), &u)
		return &u, nil
	}

	log.Println("[CACHE MISS] Fetching from DB...")

	// 🔹 DB 조회에서 age 포함
	row := s.db.QueryRow(ctx,
		`SELECT id, name, email, age FROM users WHERE id=$1`,
		id,
	)

	var u User
	err = row.Scan(&u.ID, &u.Name, &u.Email, &u.Age)
	if err != nil {
		return nil, err
	}

	// 🔹 CACHE 저장
	jsonData, _ := json.Marshal(u)
	s.rdb.Set(ctx, "user:"+id, string(jsonData), 0)

	return &u, nil
}
