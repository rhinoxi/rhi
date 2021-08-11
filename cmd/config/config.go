package config

import (
	"encoding/json"
	"os"
	"path"
)

const (
	config_folder = ".rhi"
	config_file   = "config.json"
)

func GetConfigDir() string {
	dirname, _ := os.UserHomeDir()
	return path.Join(dirname, config_folder)
}

func getConfigFilePath() string {
	return path.Join(GetConfigDir(), config_file)
}

type config struct {
	Projects []string `json:"projects"`
}

func (c *config) AddProject(s string) {
	for _, p := range c.Projects {
		if p == s {
			return
		}
	}
	c.Projects = append(c.Projects, s)
}

func (c *config) RemoveProject(s string) bool {
	for i, p := range c.Projects {
		if p == s {
			c.Projects = append(c.Projects[:i], c.Projects[i+1:]...)
			return true
		}
	}
	return false
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

func LoadConfig() (*config, error) {
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

func init() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	if err := os.MkdirAll(path.Join(dirname, config_folder), 0755); err != nil {
		panic(err)
	}
}
