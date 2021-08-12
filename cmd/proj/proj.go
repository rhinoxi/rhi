package proj

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"sort"

	"github.com/rhinoxi/rhi/cmd/config"
	"github.com/rhinoxi/rhi/cmd/initRhi"
	"github.com/spf13/cobra"

	_ "embed"
)

const (
	shell_file = "rhi_proj.sh"
	data_file  = "proj.data"
)

func getDataPath() string {
	return path.Join(config.MustGetDataDir(), data_file)
}

type ProjectData struct {
	Projects []proj `json:"projects"`
}

func (d *ProjectData) Save() error {
	f, err := os.Create(getDataPath())
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "\t")
	return enc.Encode(d)
}

func (d *ProjectData) AddProject(s string) {
	b := path.Base(s)
	suf := 1
	bsuf := b
	for _, p := range d.Projects {
		if p.Value() == s {
			return
		}
		if p.Key() == bsuf {
			bsuf = fmt.Sprintf("%s-%d", b, suf)
			suf++
		}
	}
	d.Projects = append(d.Projects, proj{bsuf, s})
}

func (d *ProjectData) RemoveProject(s string) bool {
	for i, p := range d.Projects {
		if p.Value() == s {
			d.Projects = append(d.Projects[:i], d.Projects[i+1:]...)
			return true
		}
	}
	return false
}

func LoadData() (*ProjectData, error) {
	var d ProjectData
	p := getDataPath()
	f, err := os.Open(p)
	if err != nil {
		return &d, nil
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(&d); err != nil {
		return nil, err
	}
	sort.Slice(d.Projects, func(i, j int) bool {
		return d.Projects[i][0] < d.Projects[j][0]
	})
	return &d, nil
}

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

type proj []string

func (p proj) Key() string {
	return p[0]
}

func (p proj) Value() string {
	return p[1]
}

func getConfigShellPath() string {
	return path.Join(config.MustGetConfigDir(), shell_file)
}

//go:embed shell/rhi_proj.sh
var shell []byte

func genShellFile() {
	f, err := os.Create(getConfigShellPath())
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(shell)
	fmt.Println(`Init success.

zsh:
	Add following line to ~/.zshrc
		source $HOME/.rhi/rhi_proj.sh
	Run:
		source ~/.zshrc

bash:
	Add following line to ~/.bashrc
		source $HOME/.rhi/rhi_proj.sh
	Run:
		source ~/.bashrc`)
}

func init() {
	initRhi.Register(genShellFile)
}
