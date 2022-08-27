package main

import (
	"github.com/gin-gonic/gin"
	"github.com/meian/hansei-board/go/core/template"
)

func main() {
	router := gin.Default()
	router.GET("/templates", func(ctx *gin.Context) {
		ctx.JSON(200, template.Names)
	})
	router.POST("/preview", Preview)
	router.Run(":3000")
}
