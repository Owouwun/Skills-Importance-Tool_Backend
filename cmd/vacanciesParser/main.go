package main

import (
	"fmt"
	"vacanciesParser/internal/core/logic"
	"vacanciesParser/internal/core/repository"
)

/*
	Вытягиваем все Golang-вакансии с hh.ru.

	Алгоритм работы:
	1. Тянем актуальные специализации из api.hh.ru/professional_roles.
	2. Отобираем id всех тех, что находятся в группе "Информационные технологии".
	3. Формируем и отправляем запрос к api.hh.ru/vacancies.
	4. Итеративно обрабатываем все новые вакансии, собирая с них необходимую информацию и добавляя в базу данных.
*/

func main() {
	repo := repository.NewRepository()
	l := logic.NewService(repo)

	vacancies := l.GetVacancies()

	fmt.Println(vacancies)
}
