package ascart

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	ratio     float64
	ascWidth  int
	ascHeight int
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "ascart <image>",
		Short:                 "convert image to ascii art",
		Args:                  cobra.MaximumNArgs(1),
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			im, err := readImage(args[0])
			if err != nil {
				logrus.Fatal(err)
			}
			asc, err := convert(im)
			if err != nil {
				logrus.Fatal(err)
			}
			draw(asc)
		},
	}
	cmd.Flags().Float64VarP(&ratio, "ratio", "r", 0.5, "column width / row height of the terminal")
	cmd.Flags().IntVarP(&ascWidth, "x", "x", 0, "ascii art width(columns)")
	cmd.Flags().IntVarP(&ascHeight, "y", "y", 0, "ascii art height(rows)")
	return cmd
}
