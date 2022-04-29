package proj

import (
	"fmt"
	"path"
	"regexp"
	"strings"

	"github.com/rhinoxi/rhi/constant"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func printProjects(projects []proj) {
	var longestKeyCount int
	for _, p := range projects {
		if len(p.Key()) > longestKeyCount {
			longestKeyCount = len(p.Key())
		}
	}
	if len(projects) == 0 {
		fmt.Println("no project")
	}
	var sb strings.Builder
	for _, p := range projects {
		sb.WriteString(constant.Green)
		sb.WriteString(p.Key())
		sb.WriteString(constant.ColorReset)
		sb.WriteString(": ")
		for i := 0; i < longestKeyCount-len(p.Key()); i++ {
			sb.WriteString(" ")
		}
		sb.WriteString(p.Value())
		sb.WriteString("\n")
	}
	fmt.Print(sb.String())
}

func pickProjects(projects []proj, kw string) proj {
	pattern := strings.Join(strings.Split(kw, ""), ".*")
	m, _ := regexp.Compile("(?i)" + pattern)

	// search key first
	for _, p := range projects {
		if p.Key() == kw {
			return p
		}
	}

	for _, p := range projects {
		if m.MatchString(p.Key()) {
			return p
		}
	}

	// search basename second
	for _, p := range projects {
		if m.MatchString(path.Base(p.Value())) {
			return p
		}
	}

	// search whole path
	for _, p := range projects {
		if m.MatchString(p.Value()) {
			return p
		}
	}
	return nil
}

func newShowProjects() *cobra.Command {
	return &cobra.Command{
		Use:                   "show <project>",
		Short:                 "show projects path",
		DisableFlagsInUseLine: true,
		Args:                  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			d, err := LoadData()
			if err != nil {
				logrus.Fatal(err)
			}
			projects := d.Projects
			if len(args) > 0 {
				project := pickProjects(projects, args[0])
				if project == nil {
					return
				}
				projects = []proj{project}
			}
			printProjects(projects)
		},
	}
}
