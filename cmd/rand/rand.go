package rand

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/rhinoxi/rhi/constant"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	upper         bool
	typ           string
	allowedUpper  bool
	allowedSymbol bool
)

func NewCmd() *cobra.Command {
	randCmd := &cobra.Command{
		Use:                   "rand [<len>]",
		Short:                 "Generate random string",
		Args:                  cobra.MaximumNArgs(1),
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			length := 8
			var err error
			if len(args) > 0 {
				length, err = strconv.Atoi(args[0])
				if err != nil {
					logrus.Fatalf("length should be int, got %s\n", args[0])
				}
			}
			r := randomString(length, allowedUpper, allowedSymbol)
			if upper {
				r = strings.ToUpper(r)
			}
			fmt.Println(r)
		},
	}
	randCmd.Flags().BoolVarP(&upper, "upper", "U", false, "upper case all")
	randCmd.Flags().BoolVarP(&allowedUpper, "allowed-upper", "u", false, "allowed uppercase letters")
	randCmd.Flags().BoolVarP(&allowedSymbol, "allowed-symbol", "s", false, "allowed symbols")
	return randCmd
}

func randomString(n int, allowedUpper bool, allowedSymbol bool) string {
	letters := constant.Letters
	if allowedUpper {
		letters += constant.UpperLetters
	}
	if allowedSymbol {
		letters += constant.Symbol
	}
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
