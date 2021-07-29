package rand

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const  (
	Letters = "0123456789abcdefghijklmnopqrstuvwxyz"
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Symbol = "!_@#$%&"
)


var (
	upper bool
	typ string
	allowedUpper bool
	allowedSymbol bool
)

func NewRandCmd() *cobra.Command {
	randCmd := &cobra.Command{
		Use: "rand",
		Short: "Generate random string/int",
		Args: cobra.MaximumNArgs(1),
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
	letters := Letters
	if allowedUpper {
		letters += UpperLetters
	}
	if allowedSymbol {
		letters += Symbol
	}
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}