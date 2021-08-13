package jwt

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func getKeyFunc(algStr string) (keyFunc jwt.Keyfunc, err error) {
	switch strings.ToLower(algStr) {
	case "hs256", "hs384", "hs512":
		keyFunc = func(token *jwt.Token) (interface{}, error) {
			return []byte(ps.key), nil
		}
		return
	case "rs256", "rs384", "rs512":
		keyFunc = func(token *jwt.Token) (interface{}, error) {
			return readRsaPublicKey(ps.key)
		}
		return
	default:
		return nil, fmt.Errorf("invalid alg: %s\n", algStr)
	}
}

func readRsaPublicKey(fn string) (*rsa.PublicKey, error) {
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPublicKeyFromPEM(b)
}

func newTokenParser() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "parse <jwt token>",
		Short:                 "parse jwt token to json",
		Args:                  cobra.ExactArgs(1),
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			tokenStr := args[0]
			keyFunc, err := getKeyFunc(ps.algStr)
			if err != nil {
				logrus.Fatal(err)
			}

			t, err := jwt.Parse(tokenStr, keyFunc)
			if err != nil {
				logrus.Fatal(err)
			}

			b, _ := json.MarshalIndent(t, "", "\t")
			fmt.Println("\n" + string(b))
		},
	}
	cmd.Flags().StringVarP(&ps.key, "key", "k", "", "key")
	cmd.Flags().StringVar(&ps.algStr, "alg", "hs256", "alg")
	cmd.MarkFlagRequired("key")
	return cmd
}
