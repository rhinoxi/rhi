package proj

import (
	"fmt"
	"os"
	"path"

	"github.com/rhinoxi/rhi/cmd/config"
	"github.com/rhinoxi/rhi/cmd/initRhi"
	"github.com/spf13/cobra"
)

const (
	shell_file = "rhi_proj.sh"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proj",
		Short: "project",
	}
	cmd.AddCommand(newAddProject())
	cmd.AddCommand(newShowProjects())
	cmd.AddCommand(newRemoveProject())

	return cmd
}

func getConfigShellPath() string {
	return path.Join(config.GetConfigDir(), shell_file)
}

func genShellFile() {
	shell := `
pcd() {
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

padd() {
	local folder=${1}
	[ -z $folder ] && folder="."
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

prm() {
	local folder=${1}
	[ -z $folder ] && return
	rhi proj rm $folder
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
	initRhi.Register(genShellFile)
}
