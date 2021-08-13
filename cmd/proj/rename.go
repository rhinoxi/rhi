package proj

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newRenameProject() *cobra.Command {
	return &cobra.Command{
		Use:                   "rename <old> <new>",
		Short:                 "rename project short name",
		DisableFlagsInUseLine: true,
		Args:                  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			d, err := LoadData()
			if err != nil {
				logrus.Fatal(err)
			}
			if err := d.Rename(args[0], args[1]); err != nil {
				logrus.Fatal(err)
			} else {
				fmt.Printf("success: %s -> %s\n", args[0], args[1])
			}
			d.Save()
		},
	}
}
