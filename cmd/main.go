package main

import (
	"log"

	"github.com/rommms07/idream-erp/api"
	"github.com/rommms07/idream-erp/core/auth/facebook"
)

func main() {
	go func() {
		token, err := facebook.Login(&facebook.FacebookLoginOptions{})
		if err != nil {
			log.Fatalf(err.Error())
			return
		}

		println(token.Access_token)
	}()
	api.Start()
}
