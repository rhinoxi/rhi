package cmd

import (
	"github.com/rhinoxi/rhi/cmd/cli"
	"github.com/rhinoxi/rhi/cmd/rand"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "rhi",
	Short: "rhi is the cli kits for rhinoxi",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(rand.NewRandCmd())
	rootCmd.AddCommand(cli.NewCliCmd())
}