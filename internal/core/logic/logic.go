package logic

type Repository interface {
	GetVacancies() []Vacancy
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetVacancies() []Vacancy {
	return s.repo.GetVacancies()
}
