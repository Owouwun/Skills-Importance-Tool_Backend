package redis

import (
	"log"
	"time"
)

func GetLastExecutionDate() time.Time {
	LEDate, err := Rdb.Get(Ctx, "LastExecutionDate").Result()
	// Возможно, мы запускаем программу впервые. Либо слишком долго не запускали программу и ключ протух.
	if err != nil {
		log.Printf("Redis: Ошибка определения даты последнего запуска скрипта: %v\n", err)
		return time.UnixMilli(0) // Будем собирать вакансии за всё время. Можно возвращать ошибку, но это усложнит код.
	}

	t, err := time.Parse(time.RFC3339, LEDate)
	if err != nil {
		log.Printf("Ошибка парсинга последней даты последнего запуска скрипта: %v\n", err)
		return time.UnixMilli(0)
	}

	return t
}

func UpdateLastExecutionDate() {
	Rdb.Set(Ctx, "LastExecutionDate", time.Now().Format(time.RFC3339), 0)
}
