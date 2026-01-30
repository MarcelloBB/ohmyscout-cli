package theme

import "context"

type Repository interface {
	ListThemes(ctx context.Context) ([]string, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ListThemes(ctx context.Context) ([]string, error) {
	return s.repo.ListThemes(ctx)
}
