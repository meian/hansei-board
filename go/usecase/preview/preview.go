package preview

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/meian/hansei-board/go/components/draw"
	"github.com/meian/hansei-board/go/core/font"
	"github.com/meian/hansei-board/go/core/generics"
	"github.com/meian/hansei-board/go/core/template"
	"github.com/pkg/errors"
	xdraw "golang.org/x/image/draw"
)

type Param struct {
	Template string
	Text     string
	FontType font.Type
	Color    string
	Scale    float64
}

func (p Param) Validate(ctx context.Context) error {
	return validation.ValidateStructWithContext(ctx, &p,
		validation.Field(&p.Template,
			validation.Required,
			validation.In(generics.ToAnySlice(template.Names)...)),
		validation.Field(&p.Text,
			validation.Required),
		validation.Field(&p.FontType,
			validation.Required,
			validation.In(generics.ToAnySlice(font.TypeValues())...)),
		validation.Field(&p.Color,
			validation.Required,
			validation.Match(regexp.MustCompile("^[0-9a-fA-F]{6}$"))),
		validation.Field(&p.Scale,
			validation.Required,
			validation.Min(0.1),
			validation.Max(0.999)),
	)
}

func (p Param) ToDrawParam() draw.Param {
	c := color.NRGBA{A: 255}
	_, _ = fmt.Sscanf(p.Color, "%02x%02x%02x", &c.R, &c.G, &c.B)
	return draw.Param{
		Template: p.Template,
		Text:     p.Text,
		FontType: p.FontType,
		Size:     20,
		Color:    c,
	}
}

func Exec(ctx context.Context, p Param) (image.Image, error) {
	if err := p.Validate(ctx); err != nil {
		return nil, errors.Wrap(err, "invalid request")
	}
	img, err := draw.Image(p.ToDrawParam())
	if err != nil {
		return nil, errors.Wrap(err, "failed at draw")
	}
	res := image.NewRGBA(scaledRect(img, p.Scale))
	xdraw.CatmullRom.Scale(res, res.Bounds(), img, img.Bounds(), xdraw.Over, nil)
	return res, nil
}

func scaledRect(img image.Image, scale float64) image.Rectangle {
	b := img.Bounds()
	w := scale * float64(b.Dx())
	h := scale * float64(b.Dy())
	return image.Rect(0, 0, int(w), int(h))
}
