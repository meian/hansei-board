package main

import (
	"bytes"
	"image/jpeg"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/meian/hansei-board/go/core/font"
	"github.com/meian/hansei-board/go/usecase/preview"
)

type PreviewRequest struct {
	Template string `json:"template" binding:"required,min=1"`
	Text     string `json:"text" binding:"required"`
	Font     int    `json:"font" binding:"required,min=1,max=2"`
	Color    string `json:"color" binding:"required,len=6"`
}

func (r PreviewRequest) ToUsecaseParam() preview.Param {
	return preview.Param{
		Template: r.Template,
		Text:     r.Text,
		FontType: font.Type(r.Font),
		Color:    r.Color,
	}
}

func Preview(ctx *gin.Context) {
	var req PreviewRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Print(err)
		ctx.String(400, "invalid body")
		ctx.Abort()
		return
	}
	img, err := preview.Exec(ctx, req.ToUsecaseParam())
	if err != nil {
		log.Print(err)
		ctx.String(400, "cannot preview")
		ctx.Abort()
		return
	}
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: 90}); err != nil {
		log.Print(err)
		ctx.String(400, "cannot encode to jpeg")
		return
	}
	ctx.Data(200, "image/jpeg", buf.Bytes())
}
