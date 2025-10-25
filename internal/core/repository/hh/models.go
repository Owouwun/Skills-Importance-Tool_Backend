package hh

import (
	"vacanciesParser/internal/core/logic"
	"vacanciesParser/internal/core/repository/hh/api"
)

type VacanciesResponse struct {
	api.VacanciesResponse
}

func NewVacanciesResponse(vr api.VacanciesResponse) VacanciesResponse {
	return VacanciesResponse{vr}
}

func (vs VacanciesResponse) ToLogic() []logic.Vacancy {
	logicVacancies := make([]logic.Vacancy, 0)

	for _, v := range vs.Items {
		logicVacancies = append(logicVacancies, logic.Vacancy{
			Name:        v.Name,
			Salary:      logic.VacancySalary(v.Salary),
			PublishedAt: v.PublishedAt,
			URL:         v.URL,
			Employer:    logic.VacancyEmployer(v.Employer),
		})
	}

	return logicVacancies
}
