package effect

import (
	"context"
	"image"
	"image/color"

	"github.com/disintegration/imaging"
)

type rotate struct {
	effect
}

func (r *rotate) Transform(ctx context.Context, img image.Image, params map[string]interface{}) (image.Image, error) {
	angle, err := floatBinder("angle", params)

	if err != nil {
		return nil, err
	}

	bgColor, err := colorBinder("bgcolor", params)

	if err != nil {
		bgColor = color.Transparent
	}

	img = imaging.Rotate(img, angle, bgColor)
	return img, ctx.Err()
}
