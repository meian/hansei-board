package draw

import (
	"image"
	"image/color"
	"image/draw"

	gofont "golang.org/x/image/font"
	"golang.org/x/image/math/fixed"

	"github.com/meian/hansei-board/go/components/text"
	"github.com/meian/hansei-board/go/core/font"
	"github.com/meian/hansei-board/go/core/template"
	"github.com/pkg/errors"
)

var ErrTooLongText = errors.New("too long text")

type Param struct {
	Template string
	Text     string
	FontType font.Type
	Size     int
	Color    color.Color
}

// Image は p の内容に従って描写した画像を取得する
func Image(p Param) (image.Image, error) {
	tpl, err := template.ByName(p.Template)
	if err != nil {
		return nil, errors.Wrapf(err, "failed at template: %s", p.Template)
	}
	src, err := tpl.Image()
	if err != nil {
		return nil, errors.Wrapf(err, "failed at template image: %s", p.Template)
	}

	if len(p.Text) == 0 {
		return src, nil
	}

	face, err := font.LoadFace(p.FontType, p.Size)
	if err != nil {
		return nil, errors.Wrapf(err, "failed at font face: type=%v, size=%d", p.FontType, p.Size)
	}

	lines := text.LinesByWidth(face, p.Text, tpl.Board.TextWidth())
	if len(lines) > tpl.Board.MaxLines {
		return nil, ErrTooLongText
	}

	dst := image.NewRGBA(src.Bounds())
	draw.Draw(dst, dst.Bounds(), src, image.Point{}, draw.Src)

	dr := &gofont.Drawer{
		Dst:  dst,
		Src:  image.NewUniform(p.Color),
		Face: face,
	}
	drawTexts(dr, tpl.Board, lines)

	return dst, nil
}

// drawTexts は画像に文字列リストを描き込む
func drawTexts(dr *gofont.Drawer, board template.BoardInfo, lines []string) {
	bx := fixed.I(board.TextLeft())
	dr.Dot.Y = fixed.I(board.TextTop())
	for _, line := range lines {
		dr.Dot.X = bx
		dr.DrawString(line)
		dr.Dot.Y += fixed.I(board.Height)
	}
}
