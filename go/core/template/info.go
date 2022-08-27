package template

import (
	"bytes"
	"image"
	_ "image/jpeg"

	"github.com/meian/hansei-board/go/assets"
	"github.com/pkg/errors"
)

var images = map[string]image.Image{}

type Info struct {
	Name     string
	FileName string
	Board    BoardInfo
}

// Image はテンプレートの Image を取得する
func (i Info) Image() (image.Image, error) {
	if img, ok := images[i.FileName]; ok {
		return img, nil
	}
	ib, err := assets.Img(i.FileName)
	if err != nil {
		return nil, errors.Wrapf(err, "failed at load image binary: %s", i.Name)
	}
	img, _, err := image.Decode(bytes.NewReader(ib))
	if err != nil {
		return nil, errors.Wrapf(err, "failed at decode: %s", i.Name)
	}
	images[i.FileName] = img
	return img, nil
}

type BoardInfo struct {
	Bounds     image.Rectangle
	SideMargin int
	Height     int
	TopMargin  int
	MaxLines   int
}

func (bi BoardInfo) Left() int {
	return bi.Bounds.Min.X
}

func (bi BoardInfo) Right() int {
	return bi.Bounds.Max.X
}

func (bi BoardInfo) Top() int {
	return bi.Bounds.Min.Y
}

func (bi BoardInfo) Bottom() int {
	return bi.Bounds.Max.Y
}

func (bi BoardInfo) Origin() image.Point {
	return bi.Bounds.Min
}

func (bi BoardInfo) TextLeft() int {
	return bi.Left() + bi.SideMargin
}

func (bi BoardInfo) TextRight() int {
	return bi.Right() - bi.SideMargin
}

func (bi BoardInfo) TextTop() int {
	return bi.Top() + bi.TopMargin
}

func (bi BoardInfo) TextWidth() int {
	return bi.Right() - bi.Left() - (bi.SideMargin * 2)
}
