package logic

type Repository interface {
	GetVacancies()
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetVacancies() {
	s.repo.GetVacancies()
}
