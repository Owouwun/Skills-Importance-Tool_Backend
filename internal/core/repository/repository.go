package repository

import (
	"context"
	"log"
	"vacanciesParser/internal/core/repository/mongodb"
	"vacanciesParser/internal/core/repository/readonly/hh"
	"vacanciesParser/internal/core/repository/redis"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (repo *Repository) GetVacancies() {
	from := redis.GetLastExecutionDate()

	// Лучше сразу обновлять дату, чтобы не потерять вакансии, которые появятся в процессе работы программы.
	redis.UpdateLastExecutionDate()

	vacancies := make([]mongodb.Vacancy, 0)
	vacancies = append(vacancies, hh.GetVacancies(from).ToMongo()...)

	if len(vacancies) == 0 {
		log.Printf("Новые вакансии на hh.ru не найдены.")
		return
	}

	docs := make([]interface{}, len(vacancies))
	for i, data := range vacancies {
		docs[i] = data
	}

	mongodb.InsertData(context.TODO(), docs)
}
