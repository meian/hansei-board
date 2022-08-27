package template

import (
	"image"

	"github.com/pkg/errors"
)

var (
	templates = map[string]Info{}
	Names     = []string{}

	Cat1 = Info{
		Name:     "cat1",
		FileName: "cat1.jpg",
		Board: BoardInfo{
			Bounds:     image.Rect(165, 260, 750, 579),
			SideMargin: 10,
			Height:     64,
			TopMargin:  48,
			MaxLines:   5,
		},
	}
)

func init() {
	addPhotoInfo(Cat1)
}

func addPhotoInfo(pi Info) {
	templates[pi.Name] = pi
	Names = append(Names, pi.Name)
}

// ByName は name で指定したテンプレートを取得する
func ByName(name string) (*Info, error) {
	info, ok := templates[name]
	if !ok {
		return nil, errors.Errorf("not found template: %s", name)
	}
	return &info, nil
}
