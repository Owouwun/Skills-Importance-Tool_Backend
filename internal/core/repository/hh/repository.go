package hh

import (
	"fmt"
	"vacanciesParser/internal/core/repository/hh/api"
	"vacanciesParser/internal/core/repository/hh/cache"
)

func GetITRolesIDs() []string {
	ans := cache.GetITRolesIDs()
	if len(ans) > 0 {
		fmt.Printf("Restored from cache: %v", ans)
		return ans
	}

	ids := api.GetITRolesIDs()

	cache.SetITRolesIDs(ids)

	return ids
}

func BuildQuery(roles []string) string {
	return api.BuildQuery(roles)
}

func GetVacancies() VacanciesResponse {
	query := BuildQuery(
		GetITRolesIDs(),
	)

	return NewVacanciesResponse(
		api.GetVacancies(query),
	)
}
