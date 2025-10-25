package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Rdb *redis.Client

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Не удалось подключиться к Redis: %v", err)
	}

	fmt.Println("Успешно подключено к Redis:", pong)
}

func GetLastExecutionDate() string {
	LEDate, err := Rdb.Get(Ctx, "LastExecutionDate").Result()
	if err != nil {
		log.Printf("Redis: Ошибка определения даты последнего запуска скрипта: %v", err)
	}

	return LEDate
}

func UpdateExecutionDate() {
	Rdb.Set(Ctx, "LastExecutionDate", time.Now().String(), 0)
}
