package proj

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func printProjects(projects []string) {
	for _, folder := range projects {
		fmt.Println(folder)
	}
}

func newShowProjects() *cobra.Command {
	return &cobra.Command{
		Use:                   "show <project>",
		Short:                 "show projects path",
		DisableFlagsInUseLine: true,
		Args:                  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c, err := loadConfig()
			if err != nil {
				logrus.Fatal(err)
			}
			projects := c.Projects
			if len(args) > 0 {
				project := c.PickProjects(args[0])
				if project == "" {
					return
				}
				projects = []string{project}
			}
			printProjects(projects)
		},
	}
}
