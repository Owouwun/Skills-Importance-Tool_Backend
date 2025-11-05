package hh

import (
	"fmt"
	"time"
	"vacanciesParser/internal/core/repository/readonly/hh/api"
	"vacanciesParser/internal/core/repository/readonly/hh/cache"
)

func GetITRolesIDs() []string {
	ans := cache.GetITRolesIDs()
	if len(ans) > 0 {
		fmt.Printf("Restored from cache: %v\n", ans)
		return ans
	}

	ids := api.GetITRolesIDs()

	cache.SetITRolesIDs(ids)

	return ids
}

func BuildQuery(roles []string, from time.Time) string {
	return api.BuildQuery(roles, from)
}

func GetVacancies(from time.Time) api.Vacancies {
	vacancies := make([]api.Vacancy, 0)

	query := BuildQuery(
		GetITRolesIDs(),
		from,
	)

	pages := 1
	for page := 1; page <= pages; page++ {
		q := api.UpdateQueryPage(query, page)

		vs := api.GetVacancies(q)

		vacancies = append(vacancies, vs.Items...)

		// Лучше обновлять количество страниц на случай, если в процессе работы новые вакансии породят новую страницу.
		// Можно оптимизировать, обновляя pages только когда page его догоняет (возможно даже запустив отдельный for), но выглядит излишним.
		pages = vs.Pages
	}

	return vacancies
}
