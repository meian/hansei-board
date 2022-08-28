package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/meian/hansei-board/go/core/template"
)

var (
	AConf AppConfig
	PConv PreviewConfig
)

func main() {
	envconfig.Process("APP", &AConf)
	log.Printf("app config: %+v", AConf)
	envconfig.Process("APP_PREVIEW", &PConv)
	log.Printf("preview config: %+v", PConv)
	router := gin.Default()
	router.GET("/templates", func(ctx *gin.Context) {
		ctx.JSON(200, template.Names)
	})
	router.POST("/preview", Preview)
	router.Run(fmt.Sprintf(":%d", AConf.Port))
}
