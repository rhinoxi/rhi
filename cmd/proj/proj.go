package proj

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

const (
	config_folder = ".rhi"
	config_file   = "config.json"
	shell_file    = "rhi_proj.sh"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proj",
		Short: "project",
	}
	cmd.AddCommand(newAddProject())
	cmd.AddCommand(newShowProjects())

	return cmd
}

func getConfigPath() string {
	dirname, _ := os.UserHomeDir()
	return path.Join(dirname, config_folder)
}

func getConfigFilePath() string {
	return path.Join(getConfigPath(), config_file)
}

func getConfigShellPath() string {
	return path.Join(getConfigPath(), shell_file)
}

func GenShellFile() {
	shell := `
p() {
	local output="$(rhi proj show ${1})"
	if [ -z "${1}" ]
	then
		echo $output
	else
		if [ ! -z $output ]
		then
			echo $output
			cd ${output}
		fi
	fi
}

pa() {
	local folder=${1}
	[ -z $folder ] && return
	if [[ $folder != /* ]]
	then
		if [ -d $folder ]
		then
			folder=$(cd $(cd $(dirname $folder); pwd)/$(basename $folder); pwd)
		elif [ -f $folder ]
		then
			folder=$(cd $(dirname $folder); pwd)
		else
			echo $folder not exist
			return
		fi
	fi
	cd $folder
	rhi proj add $folder
}
`

	f, err := os.Create(getConfigShellPath())
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(shell)
	fmt.Println("add following line to ~/.zshrc")
	fmt.Println("\tsource $HOME/.rhi/rhi_proj.sh")
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
