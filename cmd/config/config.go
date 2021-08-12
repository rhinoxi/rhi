package config

import (
	"os"
	"path"

	"github.com/sirupsen/logrus"

	"github.com/rhinoxi/rhi/constant"
)

func MustGetConfigDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		logrus.Fatal(err)
	}
	return path.Join(dirname, constant.ConfigFolder)
}

func MustGetDataDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		logrus.Fatal(err)
	}
	return path.Join(dirname, constant.DataFolder)
}

func init() {
	if err := os.MkdirAll(MustGetDataDir(), 0755); err != nil {
		panic(err)
	}
}
