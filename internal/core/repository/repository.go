package repository

import (
	"vacanciesParser/internal/core/logic"
	"vacanciesParser/internal/core/repository/hh"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (repo *Repository) GetVacancies() []logic.Vacancy {
	vacancies := make([]logic.Vacancy, 0)

	vacancies = append(vacancies, hh.GetVacancies().ToLogic()...)

	return vacancies
}
