package initRhi

import (
	"github.com/rhinoxi/rhi/cmd/proj"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	return &cobra.Command{
		Use:                   "init",
		Short:                 "initial rhi",
		DisableFlagsInUseLine: true,
		Args:                  cobra.MaximumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			proj.GenShellFile()
		},
	}
}
