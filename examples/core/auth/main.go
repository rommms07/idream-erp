package main

import (
	"log"
	"strings"

	"github.com/rommms07/idream-erp/api"
	"github.com/rommms07/idream-erp/core/auth/facebook"
)

func main() {

	go func() {
		opts := &facebook.FacebookLoginOptions{LoginType: facebook.LoginType_BUSINESS}

		println(facebook.LoginUrl(opts))

		token, err := facebook.Login(opts)
		if err != nil {
			log.Fatalf(err.Error())
			return
		}

		ltoken, _ := token.GetLongLivedToken(opts.LoginType)
		println(ltoken.Access_token)
		println(strings.Repeat("=", 25))
	}()

	api.Start()
}
