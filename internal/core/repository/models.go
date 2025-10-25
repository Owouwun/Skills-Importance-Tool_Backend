package repository

import (
	"vacanciesParser/internal/core/repository/hh"
)

type Vacancies struct {
	hh.VacanciesResponse
}

func NewVacancies(vr hh.VacanciesResponse) Vacancies {
	return Vacancies{vr}
}
