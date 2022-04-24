package dl

import (
	"github.com/spf13/cobra"
)

var (
	url         string
	selector    string
	outdir      string
	limit       int
	concurrency int
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "dl",
		Short:                 "downloader",
		DisableFlagsInUseLine: true,
	}

	cmd.PersistentFlags().StringVarP(&url, "url", "u", "", "url")
	cmd.PersistentFlags().StringVarP(&selector, "selector", "s", "", "selector")
	cmd.PersistentFlags().StringVarP(&outdir, "out", "o", "", "outdir")
	cmd.PersistentFlags().IntVarP(&limit, "limit", "l", 0, "limit")
	cmd.PersistentFlags().IntVarP(&concurrency, "concurrency", "c", 1, "concurrency")

	cmd.MarkPersistentFlagRequired("url")
	cmd.MarkPersistentFlagRequired("selector")
	cmd.MarkPersistentFlagRequired("out")

	cmd.AddCommand(newImgDl())
	return cmd
}
