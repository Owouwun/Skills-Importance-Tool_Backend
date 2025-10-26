package repository

import (
	"vacanciesParser/internal/core/logic"
	"vacanciesParser/internal/core/repository/hh"
	"vacanciesParser/internal/core/repository/redis"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (repo *Repository) GetVacancies() []logic.Vacancy {
	vacancies := make([]logic.Vacancy, 0)
	from := redis.GetLastExecutionDate()

	// Лучше сразу обновлять дату, чтобы не потерять вакансии, которые появятся в процессе работы программы.
	redis.UpdateLastExecutionDate()

	vacancies = append(vacancies, hh.GetVacancies(from).ToLogic()...)

	return vacancies
}
