package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/rommms07/idream-erp/core/auth/facebook"
	"github.com/rommms07/idream-erp/helpers/loader"
)

func Start() (err error) {
	router := gin.New()
	config := loader.AppConfig()

	router.GET(config.FbRedirectUri, facebook.FbRedirectHandler)

	switch config.ServerProto {
	case "http":
		err = router.Run(config.ServerAddr)
	case "https":
		err = router.RunTLS(config.ServerAddr, config.ServerCertFile, config.ServerKeyFile)
	default:
		err = errors.New("error: invalid server protocol")
	}

	return
}
