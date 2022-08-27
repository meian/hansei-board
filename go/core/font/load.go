package font

import (
	"github.com/golang/freetype/truetype"
	"github.com/meian/hansei-board/go/assets"
	"github.com/pkg/errors"
	"golang.org/x/image/font"
)

var fontNames = map[Type]string{
	MonoSpaced:  "GenJyuuGothicL-Regular.ttf",
	Propotional: "GenJyuuGothicL-P-Regular.ttf",
}

var fonts = map[Type]*truetype.Font{}
var faceMap = map[Type]map[int]font.Face{}

// LoadFace は指定した種別のフォントとサイズを用いたフォントフェイスを返す
func LoadFace(t Type, size int) (font.Face, error) {
	if _, ok := fontNames[t]; !ok {
		return nil, errors.Errorf("not supported type: %v", t)
	}
	faces, ok := faceMap[t]
	if !ok {
		faces = map[int]font.Face{}
		faceMap[t] = faces
	}
	if face, ok := faces[size]; ok {
		return face, nil
	}
	font, err := loadFont(t)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	face := truetype.NewFace(font, &truetype.Options{Size: float64(size)})
	faces[size] = face
	return face, nil
}

func loadFont(t Type) (*truetype.Font, error) {
	if font, ok := fonts[t]; ok {
		return font, nil
	}
	fb, err := assets.Font(fontNames[t])
	if err != nil {
		return nil, errors.Wrapf(err, "not supported type: %v", t)
	}
	font, err := truetype.Parse(fb)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid font file: %v", t)
	}
	return font, nil
}
