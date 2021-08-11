package proj

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type config struct {
	Projects []string `json:"projects"`
}

func (c *config) addProject(s string) {
	for _, p := range c.Projects {
		if p == s {
			return
		}
	}
	c.Projects = append(c.Projects, s)
}

func (c *config) Save() error {
	f, err := os.Create(getConfigFilePath())
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "\t")
	return enc.Encode(c)
}

func (c *config) PickProjects(kw string) string {
	pattern := strings.Join(strings.Split(kw, ""), ".*")
	m, _ := regexp.Compile(pattern)

	for _, p := range c.Projects {
		if m.MatchString(p) {
			return p
		}
	}
	return ""
}

func loadConfig() (*config, error) {
	var c config
	p := getConfigFilePath()
	f, err := os.Open(p)
	if err != nil {
		return &c, nil
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

func newAddProject() *cobra.Command {
	return &cobra.Command{
		Use:                   "add <path>",
		Short:                 "add project path",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			// load config
			c, err := loadConfig()
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
				c.addProject(folder)
			}
			c.Save()
		},
	}
}
