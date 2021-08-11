package cs

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

const (
	indentCharacter = "    "
	Red             = "\033[31m"
	Green           = "\033[32m"
	Yellow          = "\033[33m"
	Blue            = "\033[34m"
	Purple          = "\033[35m"
	Cyan            = "\033[36m"
)

var colors = []string{Green, Yellow, Blue, Purple, Cyan}

func getCsKeys() []string {
	keys := make([]string, 0, len(csm))
	for key := range csm {
		keys = append(keys, key)
	}
	return keys
}

func parenKeys(keys []string) string {
	p := make([]string, len(keys))
	for i, k := range keys {
		p[i] = "[" + k + "]"
	}
	return strings.Join(p, " ")
}

func pickColor(n int) string {
	return colors[n%len(colors)]
}

type cmdFormatter struct {
	sb strings.Builder
}

func (cf *cmdFormatter) add(key string, value interface{}, indent int) {
	cf.sb.WriteString(pickColor(indent))
	cf.sb.WriteString(key)
	cf.sb.WriteString("\033[0m")
	cf.sb.WriteString("\n")
	for i := 0; i <= indent; i++ {
		cf.sb.WriteString(indentCharacter)
	}
	switch tv := value.(type) {
	case string:
		cf.sb.WriteString(tv)
		cf.sb.WriteString("\n")
	case csType:
		for k, v := range tv {
			cf.add(k, v, indent+1)
		}
	}
}

func (cf cmdFormatter) String() string {
	return cf.sb.String()
}

func NewCmd() *cobra.Command {
	csKeys := getCsKeys()
	return &cobra.Command{
		Use:                   "cs " + parenKeys(csKeys),
		Short:                 "show cheatsheet",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			var cf cmdFormatter
			if len(args) == 0 {
				args = csKeys
			}
			for _, arg := range args {
				if line := csm[arg]; line != nil {
					cf.add(arg, line, 0)
				}
			}
			fmt.Println(cf)
		},
		ValidArgs: csKeys,
	}
}
