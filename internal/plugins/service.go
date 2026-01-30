package plugin

import "context"

type Repository interface {
	ListPlugins(ctx context.Context) ([]string, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ListPlugins(ctx context.Context) ([]string, error) {
	return s.repo.ListPlugins(ctx)
}
