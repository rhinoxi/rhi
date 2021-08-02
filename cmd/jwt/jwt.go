package jwt

import "github.com/spf13/cobra"

func NewJwtCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "jwt",
		Short: "jwt token generator/parser",
		DisableFlagsInUseLine: true,
	}
	cmd.AddCommand(newTokenGenerator())
	cmd.AddCommand(newTokenParser())
	return cmd
}