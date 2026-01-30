package cmd

import (
	"context"
	"fmt"
	//	"os"

	"github.com/spf13/cobra"

	"github.com/MarcelloBB/ohmyscout-cli/internal/github"
	"github.com/MarcelloBB/ohmyscout-cli/internal/theme"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Adiciona ou modifica configurações do ohmyzsh",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		githubClient := github.NewClient()
		repo := github.NewRepository(githubClient)
		service := theme.NewService(repo)

		themes, err := service.ListThemes(ctx)
		if err != nil {
			return err
		}

		for _, theme := range themes {
			fmt.Println(theme)
		}

		return nil
	},
}
