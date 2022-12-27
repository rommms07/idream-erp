package main

import (
	"log"

	"github.com/rommms07/idream-erp/api"
	"github.com/rommms07/idream-erp/core/auth/facebook"
)

func main() {
	go func() {
		opts := &facebook.FacebookLoginOptions{}

		println(facebook.LoginUrl(opts))

		token, err := facebook.Login(opts)
		if err != nil {
			log.Fatalf(err.Error())
			return
		}

		ltoken, _ := token.GetLongLivedToken()
		println(ltoken.Access_token)
	}()
	api.Start()
}
