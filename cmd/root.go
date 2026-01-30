package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "ohmyscout",
	Short: "tool for setting up ohmyzsh configurations",
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}
