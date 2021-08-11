package proj

import (
	"fmt"
	"path"
	"regexp"
	"strings"

	"github.com/rhinoxi/rhi/cmd/config"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func printProjects(projects []string) {
	if len(projects) == 0 {
		fmt.Println("no project")
	}
	for _, folder := range projects {
		fmt.Println(folder)
	}
}

func pickProjects(projects []string, kw string) string {
	pattern := strings.Join(strings.Split(kw, ""), ".*")
	m, _ := regexp.Compile(pattern)

	// search basename first
	for _, p := range projects {
		base := path.Base(p)
		if m.MatchString(base) {
			return p
		}
	}

	// search whole path
	for _, p := range projects {
		p = path.Base(p)
		if m.MatchString(p) {
			return p
		}
	}
	return ""
}

func newShowProjects() *cobra.Command {
	return &cobra.Command{
		Use:                   "show <project>",
		Short:                 "show projects path",
		DisableFlagsInUseLine: true,
		Args:                  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c, err := config.LoadConfig()
			if err != nil {
				logrus.Fatal(err)
			}
			projects := c.Projects
			if len(args) > 0 {
				project := pickProjects(projects, args[0])
				if project == "" {
					return
				}
				projects = []string{project}
			}
			printProjects(projects)
		},
	}
}
