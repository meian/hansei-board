package main

import (
	"context"
	"image/jpeg"
	"os"

	"github.com/meian/hansei-board/go/core/font"
	"github.com/meian/hansei-board/go/usecase/preview"
)

func main() {

	r := preview.Param{
		Template: "cat1",
		Text:     "あいうえおかきくけこ\nさしすせそたちつてと\nなにぬねのはひふへほ\nまみむめもやゆよらりるれろわをん\n",
		FontType: font.Propotional,
		Color:    "ff0000",
	}

	dst, err := preview.Exec(context.Background(), r)
	if err != nil {
		panic(err)
	}

	// mc := color.RGBA{R: 80, G: 80, B: 80, A: 40}
	// mask := image.NewUniform(mc)
	// draw.Draw(dst, board.Bounds, mask, board.Origin(), draw.Over)

	w, err := os.Create("/workspace/assets/result.jpg")
	if err != nil {
		panic(err)
	}
	defer w.Close()
	jpeg.Encode(w, dst, &jpeg.Options{Quality: 90})
}
