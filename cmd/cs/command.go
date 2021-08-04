package cs

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func getCsKeys() []string {
	keys := make([]string, 0, len(cheatsheet))
	for key := range cheatsheet {
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

func NewCsCmd() *cobra.Command {
	csKeys := getCsKeys()
	return &cobra.Command{
		Use:                   "cs " + parenKeys(csKeys),
		Short:                 "show cheatsheet",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			var sb strings.Builder
			if len(args) == 0 {
				args = csKeys
			}
			for _, arg := range args {
				if line := cheatsheet[arg]; line != "" {
					sb.WriteString(arg)
					sb.WriteString("\n\t")
					sb.WriteString(line)
					sb.WriteString("\t")
				}
			}
			fmt.Println(sb.String())
		},
		ValidArgs: csKeys,
	}
}
