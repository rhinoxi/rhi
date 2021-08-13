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

type params struct {
	key       string
	algStr    string
	claimFile string
}

var (
	ps             params
	algStrToMethod = map[string]jwt.SigningMethod{
		"hs256": jwt.SigningMethodHS256,
		"hs384": jwt.SigningMethodHS384,
		"hs512": jwt.SigningMethodHS512,
		"rs256": jwt.SigningMethodRS256,
		"rs384": jwt.SigningMethodRS384,
		"rs512": jwt.SigningMethodRS512,
	}
)

type claims map[string]interface{}

func parseClaimFile(fn string) (c claims, err error) {
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &c)
	return
}

func readRsaPrivateKey(fn string) (*rsa.PrivateKey, error) {
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPrivateKeyFromPEM(b)
}

func getSignMethodAndKey(algStr string) (jwt.SigningMethod, interface{}, error) {
	algStr = strings.ToLower(algStr)
	var alg jwt.SigningMethod
	var key interface{}
	switch algStr {
	case "hs256", "hs384", "hs512":
		alg = algStrToMethod[algStr]
		key = []byte(ps.key)
	case "rs256", "rs384", "rs512":
		alg = algStrToMethod[algStr]
		var err error
		key, err = readRsaPrivateKey(ps.key)
		if err != nil {
			return nil, nil, err
		}
	default:
		return nil, nil, fmt.Errorf("invalid alg: %s\n", ps.algStr)
	}
	return alg, key, nil
}

func sign(alg jwt.SigningMethod, key interface{}, m claims) (string, error) {
	token := jwt.NewWithClaims(alg, jwt.MapClaims(m))
	return token.SignedString(key)
}

func newTokenGenerator() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "sign -c <file> -k <key>",
		Short:                 "jwt token generator",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			alg, key, err := getSignMethodAndKey(ps.algStr)
			if err != nil {
				logrus.Fatal(err)
			}

			m, err := parseClaimFile(ps.claimFile)
			if err != nil {
				logrus.Fatal(err)
			}

			tokenStr, err := sign(alg, key, m)
			if err != nil {
				logrus.Fatal(err)
			}

			fmt.Println(tokenStr)
		},
	}
	cmd.Flags().StringVarP(&ps.claimFile, "claim", "c", "", "claim json file")
	cmd.Flags().StringVarP(&ps.key, "key", "k", "", "key string(hmac-sha)/filepath(rsa)")
	cmd.Flags().StringVar(&ps.algStr, "alg", "hs256", "alg")
	cmd.MarkFlagRequired("claim")
	cmd.MarkFlagRequired("key")
	return cmd
}
