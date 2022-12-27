package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rommms07/idream-erp/core/auth/facebook"
)

func Start() {
	router := gin.New()

	router.GET("/oauthcb/facebook", facebook.FbRedirectHandler)

	router.Run(source.AppConfig().ServerAddr)
}
