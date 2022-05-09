package cmd

import (
	"github.com/rhinoxi/rhi/cmd/ascart"
	"github.com/rhinoxi/rhi/cmd/cs"
	"github.com/rhinoxi/rhi/cmd/dl"
	"github.com/rhinoxi/rhi/cmd/initRhi"
	"github.com/rhinoxi/rhi/cmd/jwt"
	"github.com/rhinoxi/rhi/cmd/proj"
	"github.com/rhinoxi/rhi/cmd/rand"
	"github.com/rhinoxi/rhi/cmd/template"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rhi",
	Short: "rhi is the cli kits for rhinoxi",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
		DisableQuote:     true,
	})
	rootCmd.AddCommand(initRhi.NewCmd())
	rootCmd.AddCommand(rand.NewCmd())
	rootCmd.AddCommand(cs.NewCmd())
	rootCmd.AddCommand(jwt.NewCmd())
	rootCmd.AddCommand(proj.NewCmd())
	rootCmd.AddCommand(ascart.NewCmd())
	rootCmd.AddCommand(dl.NewCmd())
	rootCmd.AddCommand(template.NewCmd())
}
