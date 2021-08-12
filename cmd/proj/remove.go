package proj

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newRemoveProject() *cobra.Command {
	return &cobra.Command{
		Use:                   "rm <path>",
		Short:                 "remove project from local storage",
		DisableFlagsInUseLine: true,
		Args:                  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			d, err := LoadData()
			if err != nil {
				logrus.Fatal(err)
			}
			for _, folder := range args {
				if d.RemoveProject(folder) {
					fmt.Printf("%s has been removed\n", folder)
				}
			}
			d.Save()
		},
	}
}
