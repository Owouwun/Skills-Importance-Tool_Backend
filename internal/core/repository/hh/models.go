package hh

import (
	"vacanciesParser/internal/core/repository/hh/api"
)

type VacanciesResponse struct {
	api.VacanciesResponse
}

func NewVacanciesResponse(vr api.VacanciesResponse) VacanciesResponse {
	return VacanciesResponse{vr}
}
