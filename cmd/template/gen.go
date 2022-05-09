package template

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

//go:embed templates/*
var files embed.FS

func NewCmd() *cobra.Command {
	genCmd := &cobra.Command{
		Use:   "template",
		Short: "Generate template file",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				if err := listTemplateNames(); err != nil {
					logrus.Fatal(err)
				}
				return
			}
			for _, tn := range args {
				if err := genTemplate(tn); err != nil {
					logrus.Fatal(err)
				}
			}
		},
	}

	return genCmd
}

func listTemplateNames() error {
	templates, err := fs.ReadDir(files, "templates")
	if err != nil {
		return err
	}
	for _, template := range templates {
		fmt.Println(template.Name())
	}
	return nil
}

func genTemplate(tn string) error {
	f, err := files.Open(path.Join("templates", tn))
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("%s not exist", tn)
		}
		return err
	}
	defer f.Close()
	out, err := os.OpenFile(tn, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		if errors.Is(err, fs.ErrExist) {
			return fmt.Errorf("%s already exist", tn)
		}
		return err
	}
	defer out.Close()
	io.Copy(out, f)
	return nil
}
