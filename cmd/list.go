package cmd

import (
	"context"
	"fmt"

	// "os"

	"github.com/spf13/cobra"

	"github.com/MarcelloBB/ohmyscout-cli/internal/github"
	plugin "github.com/MarcelloBB/ohmyscout-cli/internal/plugins"
	"github.com/MarcelloBB/ohmyscout-cli/internal/theme"
)

func init() {
	listCmd.Flags().BoolP("theme", "t", false, "List themes")
	listCmd.Flags().BoolP("plugin", "p", false, "List plugins")
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista plugins/temas/ferramentas",
	Args:  cobra.NoArgs,
	RunE:  runListCmd,
}

func runListCmd(cmd *cobra.Command, args []string) error {
	listThemes, _ := cmd.Flags().GetBool("theme")
	listPlugins, _ := cmd.Flags().GetBool("plugin")

	if !listThemes && !listPlugins {
		fmt.Println("no theme or plugin flag detected")
	}

	ctx := context.Background()
	githubClient := github.NewClient()
	repo := github.NewRepository(githubClient)
	themesService := theme.NewService(repo)
	pluginsService := plugin.NewService(repo)

	if listThemes {
		themes, err := themesService.ListThemes(ctx)
		if err != nil {
			return err
		}

		for _, theme := range themes {
			fmt.Println("[THEME] ", theme)
		}
	}

	if listPlugins {
		plugins, err := pluginsService.ListPlugins(ctx)
		if err != nil {
			return err
		}

		for _, plugin := range plugins {
			fmt.Println("[PLUGIN] ", plugin)
		}
	}

	return nil
}
