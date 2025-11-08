package skilltree

import (
	"context"
)

type Repository interface {
	GetSkillTree(ctx context.Context) (*Node, error)
	CreateNode(ctx context.Context, tree *NodePath) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetSkillTree(ctx context.Context) (*Node, error) {
	return s.repo.GetSkillTree(ctx)
}

func (s *Service) CreateNode(ctx context.Context, node *NodePath) error {
	return s.repo.CreateNode(ctx, node)
}
