package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
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
	cacheKey := "user:" + id

	// ------------------------------------
	// 🔹 1) 캐시 조회
	// ------------------------------------
	cacheData, err := s.rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		log.Println("[CACHE HIT]")

		var cached User
		if jsonErr := json.Unmarshal([]byte(cacheData), &cached); jsonErr == nil {
			// 🔥 stale 캐시 감지 (age=0이면 이전 구조일 수 있음)
			if cached.Age != 0 {
				return &cached, nil
			}

			log.Println("[STALE CACHE] age=0 → DB에서 다시 조회합니다.")
		} else {
			log.Println("[CACHE ERROR] JSON parse 실패 → DB 조회로 이동")
		}
	} else {
		log.Println("[CACHE MISS]")
	}

	// ------------------------------------
	// 🔹 2) DB 조회
	// ------------------------------------
	log.Println("[DB QUERY] SELECT ... FROM users WHERE id =", id)

	row := s.db.QueryRow(ctx,
		`SELECT id, name, email, age FROM users WHERE id=$1`,
		id,
	)

	var u User
	err = row.Scan(&u.ID, &u.Name, &u.Email, &u.Age)
	if err != nil {
		log.Printf("[DB ERROR] row.Scan failed: %v\n", err)
		return nil, fmt.Errorf("db error: %w", err)
	}

	log.Printf("[DB RESULT] %+v\n", u)

	// ------------------------------------
	// 🔹 3) DB 결과 캐싱
	// ------------------------------------
	jsonData, _ := json.Marshal(u)
	s.rdb.Set(ctx, cacheKey, string(jsonData), 0)

	log.Println("[CACHE SAVE] user cached:", cacheKey)

	return &u, nil
}
