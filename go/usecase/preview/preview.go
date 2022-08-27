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
)

type Param struct {
	Template string
	Text     string
	FontType font.Type
	Color    string
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
	return img, nil
}
