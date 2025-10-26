package repository

import (
	"time"
	"vacanciesParser/internal/core/logic"
	"vacanciesParser/internal/core/repository/hh"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (repo *Repository) GetVacancies() []logic.Vacancy {
	vacancies := make([]logic.Vacancy, 0)
	from := time.Date(2025, time.October, 1, 0, 0, 0, 0, time.Local) // Нужно тянуть откуда-нибудь из БД как последнюю дату подтягивания данных.

	vacancies = append(vacancies, hh.GetVacancies(from).ToLogic()...)

	return vacancies
}
