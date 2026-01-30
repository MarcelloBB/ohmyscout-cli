package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/MarcelloBB/ohmyscout-cli/internal/github"
	"github.com/MarcelloBB/ohmyscout-cli/internal/theme"
)

var rootCmd = &cobra.Command{
	Use:   "zshscout",
	Short: "Lista temas do Oh My Zsh",
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
