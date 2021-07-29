package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var cliMap = map[string]string{
	"grpc-go": `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative <path/to/proto>`,
}

func getCliKeys() []string {
	keys := make([]string, 0, len(cliMap))
	for key := range cliMap {
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

func NewCliCmd() *cobra.Command {
	cliKeys := getCliKeys()
	return &cobra.Command{
		Use: "cli " + parenKeys(cliKeys),
		Short: "show cli usage",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			var sb strings.Builder
			for _, arg := range args {
				if line := cliMap[arg]; line != "" {
					sb.WriteString(arg)
					sb.WriteString("\n\t")
					sb.WriteString(line)
					sb.WriteString("\t")
				}
			}
			fmt.Println(sb.String())
		},
		ValidArgs: cliKeys,
	}
}
