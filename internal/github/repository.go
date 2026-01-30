package github

import "context"

const themesURL = "https://api.github.com/repos/ohmyzsh/ohmyzsh/contents/themes"

type Repository interface {
	ListThemes(ctx context.Context) ([]string, error)
}

type repository struct {
	client Client
}

func NewRepository(client Client) Repository {
	return &repository{client: client}
}

type contentItem struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (r *repository) ListThemes(ctx context.Context) ([]string, error) {
	var items []contentItem

	if err := r.client.Get(ctx, themesURL, &items); err != nil {
		return nil, err
	}

	var themes []string
	for _, item := range items {
		if item.Type == "file" {
			themes = append(themes, item.Name)
		}
	}

	return themes, nil
}
