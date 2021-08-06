package cs

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

const (
	indentCharacter = "    "
)

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

type cmdFormatter struct {
	sb strings.Builder
}

func (cf *cmdFormatter) add(key string, value interface{}, indent int) {
	cf.sb.WriteString(key)
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

func NewCsCmd() *cobra.Command {
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
