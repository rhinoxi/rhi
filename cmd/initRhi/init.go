package initRhi

import (
	"github.com/spf13/cobra"
)

type initHub struct {
	Members []func()
}

var hub initHub

func Register(f func()) {
	hub.Members = append(hub.Members, f)
}

func NewCmd() *cobra.Command {
	return &cobra.Command{
		Use:                   "init",
		Short:                 "initial rhi",
		DisableFlagsInUseLine: true,
		Args:                  cobra.MaximumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			for _, m := range hub.Members {
				m()
			}
		},
	}
}
