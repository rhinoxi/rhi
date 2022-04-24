package dl

import (
	"github.com/rhinoxi/rhi/cmd/dl/image"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newImgDl() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "img",
		Short: "download images through html5 css selector",
		Long: `example:
        rhi dl img -u https://wiki.factorio.com/Main_Page -s ".factorio-icon" -o /tmp/downloaded-img -l 10 -c 2
        `,
		Run: func(cmd *cobra.Command, args []string) {
			if err := image.Download(url, selector, outdir, limit, concurrency); err != nil {
				logrus.Error(err)
			}
		},
	}
	return cmd
}
