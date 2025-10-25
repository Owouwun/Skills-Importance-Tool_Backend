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
	return hh.GetVacancies().ToLogic()
}
