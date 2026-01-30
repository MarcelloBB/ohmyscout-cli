package github

import (
	"context"
	"fmt"
)

const baseURL = "https://api.github.com/repos/ohmyzsh/ohmyzsh/contents/"
const themesURL = baseURL + "themes"
const pluginsURL = baseURL + "plugins"

type Repository interface {
	ListThemes(ctx context.Context) ([]string, error)
	ListPlugins(ctx context.Context) ([]string, error)
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
		fmt.Println("Item: ", item)
		if item.Type == "file" {
			themes = append(themes, item.Name)
		}
	}

	return themes, nil
}

func (r *repository) ListPlugins(ctx context.Context) ([]string, error) {
	var items []contentItem

	if err := r.client.Get(ctx, pluginsURL, &items); err != nil {
		return nil, err
	}

	var plugins []string
	for _, item := range items {
		if item.Type == "dir" {
			plugins = append(plugins, item.Name)
		}
	}

	return plugins, nil
}
