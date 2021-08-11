package proj

import (
	"fmt"
	"os"
	"strings"

	"github.com/rhinoxi/rhi/cmd/config"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newAddProject() *cobra.Command {
	return &cobra.Command{
		Use:                   "add <path>",
		Short:                 "add project path",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			// load config
			c, err := config.LoadConfig()
			if err != nil {
				logrus.Fatal(err)
			}
			// if exist
			for _, folder := range args {
				if !strings.HasPrefix(folder, "/") {
					logrus.Error(fmt.Errorf("relative path is not allowed: %s", folder))
					continue
				}
				info, err := os.Stat(folder)
				if err != nil {
					logrus.Error(err)
					continue
				}
				if !info.IsDir() {
					logrus.Error(fmt.Errorf("%s is a file, not folder", folder))
					continue
				}
				c.AddProject(folder)
			}
			c.Save()
		},
	}
}
